package metrics

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sega-strn/metrics-tpl/internal/storage"
)

func MetricsHandler(storage *storage.MemStorage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        parts := strings.Split(r.URL.Path, "/")
        if len(parts) != 5 { // Проверка на корректное количество частей в URL
            http.Error(w, "Not Found", http.StatusNotFound)
            return
        }

        metricType := parts[2]
        metricName := parts[3]
        metricValueStr := parts[4]

        // var err error // Объявляем err один раз

        if metricType == "gauge" {
            value, err := strconv.ParseFloat(metricValueStr, 64) // Преобразуем строку в float64
            if err != nil {
                http.Error(w, "Invalid gauge value", http.StatusBadRequest)
                return
            }
            storage.UpdateGauge(metricName, value) // Обновляем значение gauge
        } else if metricType == "counter" {
            value, err := strconv.ParseInt(metricValueStr, 10, 64) // Преобразуем строку в int64
            if err != nil {
                http.Error(w, "Invalid counter value", http.StatusBadRequest)
                return
            }
            storage.UpdateCounter(metricName, value) // Обновляем значение counter
        } else {
            http.Error(w, "Invalid metric type", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Metric received")
    }
}