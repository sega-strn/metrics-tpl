package metrics

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
)

// Переменные для хранения метрик
var (
	pollCount   int64
	metricsData = make(map[string]float64) // Для хранения собранных метрик
)

// CollectMetrics собирает метрики и отправляет их на сервер
func CollectMetrics(serverAddr string, reportIntervalSec, pollIntervalSec int) {
	serverURL := fmt.Sprintf("http://%s/update", serverAddr)
	reportTicker := time.NewTicker(time.Duration(reportIntervalSec) * time.Second)
	pollTicker := time.NewTicker(time.Duration(pollIntervalSec) * time.Second)
	defer reportTicker.Stop()
	defer pollTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			// Собираем метрики из пакета runtime
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			// Сохраняем метрики в мапу
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
			metricsData["RandomValue"] = float64(time.Now().UnixNano() % 100)

			// Увеличиваем счетчик PollCount
			atomic.AddInt64(&pollCount, 1)

		case <-reportTicker.C:
			// Отправляем все метрики на сервер
			for name, value := range metricsData {
				sendMetric(serverURL, "gauge", name, value)
			}
			// Отправляем PollCount как counter
			sendMetric(serverURL, "counter", "PollCount", float64(atomic.LoadInt64(&pollCount)))
		}
	}
}

// sendMetric отправляет одну метрику на сервер
func sendMetric(serverURL, metricType, name string, value float64) {
	url := fmt.Sprintf("%s/%s/%s/%v", serverURL, metricType, name, value)
	resp, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Printf("Error sending metric %s: %v", name, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error sending metric %s: server returned status %d", name, resp.StatusCode)
	}
}
