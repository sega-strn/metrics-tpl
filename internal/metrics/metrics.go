package metrics

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	serverAddress  = "http://localhost:8080/update"
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second // Интервал отправки метрик
)

var (
	pollCount   int64
	metricsData = make(map[string]float64) // Для хранения собранных метрик
)

// CollectMetrics собирает метрики и отправляет их на сервер.
func CollectMetrics() {
	ticker := time.NewTicker(reportInterval) // reportInterval = 10 секунд
	defer ticker.Stop()

	for {
		// Собираем метрики из пакета runtime
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		// Сохраняем собранные метрики в мапу
		metricsData["Alloc"] = float64(memStats.Alloc)
		metricsData["BuckHashSys"] = float64(memStats.BuckHashSys)
		metricsData["Frees"] = float64(memStats.Frees)
		metricsData["GCCPUFraction"] = float64(memStats.GCCPUFraction)
		metricsData["GCSys"] = float64(memStats.GCSys)
		metricsData["HeapAlloc"] = float64(memStats.HeapAlloc)
		metricsData["HeapIdle"] = float64(memStats.HeapIdle)
		metricsData["HeapInuse"] = float64(memStats.HeapInuse)
		metricsData["HeapObjects"] = float64(memStats.HeapObjects)
		metricsData["HeapReleased"] = float64(memStats.HeapReleased)
		metricsData["HeapSys"] = float64(memStats.HeapSys)
		metricsData["LastGC"] = float64(memStats.LastGC)
		metricsData["Lookups"] = float64(memStats.Lookups)
		metricsData["MCacheInuse"] = float64(memStats.MCacheInuse)
		metricsData["MCacheSys"] = float64(memStats.MCacheSys)
		metricsData["MSpanInuse"] = float64(memStats.MSpanInuse)
		metricsData["MSpanSys"] = float64(memStats.MSpanSys)
		metricsData["Mallocs"] = float64(memStats.Mallocs)
		metricsData["NextGC"] = float64(memStats.NextGC)
		metricsData["NumForcedGC"] = float64(memStats.NumForcedGC)
		metricsData["NumGC"] = float64(memStats.NumGC)
		metricsData["OtherSys"] = float64(memStats.OtherSys)
		metricsData["PauseTotalNs"] = float64(memStats.PauseTotalNs)
		metricsData["StackInuse"] = float64(memStats.StackInuse)
		metricsData["StackSys"] = float64(memStats.StackSys)
		metricsData["Sys"] = float64(memStats.Sys)
		metricsData["TotalAlloc"] = float64(memStats.TotalAlloc)

		// Увеличиваем счетчик PollCount
		atomic.AddInt64(&pollCount, 1)

		// Добавляем произвольное значение
		metricsData["RandomValue"] = float64(time.Now().UnixNano() % 100)

		// Отправка метрик на сервер по таймеру
		select {
		case <-ticker.C:
			sendAllMetrics()
		}

		time.Sleep(pollInterval) // Пауза перед следующим сбором метрик
	}
}

// sendAllMetrics отправляет все собранные метрики на сервер.
func sendAllMetrics() {
	for name, value := range metricsData {
		sendMetric("gauge", name, value) // Отправляем все gauge метрики
	}
	// Отправляем counter метрику PollCount
	sendMetric("counter", "PollCount", float64(atomic.LoadInt64(&pollCount)))
}

// sendMetric отправляет одну метрику на сервер.
func sendMetric(metricType, name string, value float64) {
	url := fmt.Sprintf("%s/%s/%s/%f", serverAddress, metricType, name, value)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
