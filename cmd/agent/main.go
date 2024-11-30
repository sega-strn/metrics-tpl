package main

import "metrics-tpl/internal/metrics"

func main() {
	go metrics.CollectMetrics()

	// Бесконечный цикл для поддержания работы агента
	select {}
}
