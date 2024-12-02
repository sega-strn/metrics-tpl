package main

import (
	"flag"
	"fmt"
	"log"
	"metrics-tpl/internal/metrics"
	"os"
	"strconv"
)

func getEnvOrDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// определим флаги командной строки
	serverAdd := flag.String("a", "localhost:8080", "HTTP server address")
	reportInterval := flag.Int("r", 10, "Report interval in seconds")
	pollInterval := flag.Int("p", 2, "Poll interval in seconds")

	// Парсим флаги
	flag.Parse()

	// Проверяем нет ли не известных флагов
	if flag.NArg() > 0 {
		log.Fatal("Unknown argument provided: " + fmt.Sprint(flag.Args()))
	}

	// Приоритет отдаётся переменным окружения
	address := getEnvOrDefault("ADDRESS", *serverAdd)

	// Получаем значения интервалов из переменных окружения
	reportEnv := getEnvOrDefault("REPORT_INTERVAL", strconv.Itoa(*reportInterval))
	pollEnv := getEnvOrDefault("POLL_INTERVAL", strconv.Itoa(*pollInterval))

	// Преобразуем строковые значения в целые числа
	reportIntervalValue, err := strconv.Atoi(reportEnv)
	if err != nil {
		log.Fatal("Invalid REPORT_INTERVAL value")
	}

	pollIntervalValue, err := strconv.Atoi(pollEnv)
	if err != nil {
		log.Fatal("Invalid POLL_INTERVAL value")
	}

	// Запускаем сбор метрик с настроенными параметрами
	go metrics.CollectMetrics(address, reportIntervalValue, pollIntervalValue)

	// Бесконечный цикл для поддержания работы
	select {}

}
