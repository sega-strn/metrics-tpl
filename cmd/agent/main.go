package main

import (
	"flag"
	"fmt"
	"log"
	"metrics-tpl/internal/metrics"
)

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

	//  Запускаем сбор метрики с настроенными параметрами
	go metrics.CollectMetrics(*serverAdd, *reportInterval, *pollInterval)

	//  Бесконечный цикл для поддержание работы
	select {}

}
