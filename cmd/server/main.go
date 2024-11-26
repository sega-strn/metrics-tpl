package main

import (
	"fmt"
	"net/http"

	"github.com/sega-strn/metrics-tpl/internal/metrics"
	"github.com/sega-strn/metrics-tpl/internal/storage"
)

func main() {

	storage := storage.NewMemStorage()

	http.HandleFunc("/update/", metrics.MetricsHandler(storage)) // Используем правильный обработчик

	fmt.Println("Server is listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
