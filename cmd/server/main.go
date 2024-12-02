package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"metrics-tpl/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Определяем флаги командной строки
	serverAddr := flag.String("a", "localhost:8081", "HTTP server address")

	// Парсим флаги
	flag.Parse()

	// Проверяем, нет ли неизвестных флагов
	if flag.NArg() > 0 {
		log.Fatalf("Unknown arguments provided: %v", flag.Args())
	}

	// Создаем хранилище
	memStorage := storage.NewMemStorage()

	// Инициализация Gin
	r := gin.Default()

	// Обработчик запроса на получение значения метрики по типу и имени
	r.GET("/value/:metricType/:metricName", func(c *gin.Context) {
		metricType := c.Param("metricType")
		metricName := c.Param("metricName")

		var value interface{}
		var exists bool

		// Проверяем, какой тип метрики запрашивается
		if metricType == "gauge" {
			value, exists = memStorage.GetGauge(metricName)
		} else if metricType == "counter" {
			value, exists = memStorage.GetCounter(metricName)
		} else {
			c.String(http.StatusBadRequest, "Invalid metric type")
			return
		}

		// Если метрика не найдена, возвращаем 404
		if !exists {
			c.String(http.StatusNotFound, "Metric not found")
			return
		}

		// Возвращаем значение как текст
		c.String(http.StatusOK, fmt.Sprintf("%v", value))
	})

	// Обработчик для сохранения метрик
	r.POST("/update/:metricType/:metricName/:metricValue", func(c *gin.Context) {
		metricType := c.Param("metricType")
		metricName := c.Param("metricName")
		metricValue := c.Param("metricValue")

		switch metricType {
		case "gauge":
			value, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid gauge value")
				return
			}
			memStorage.UpdateGauge(metricName, value)

		case "counter":
			value, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid counter value")
				return
			}
			memStorage.UpdateCounter(metricName, value)

		default:
			c.String(http.StatusBadRequest, "Invalid metric type")
			return
		}

		c.Status(http.StatusOK)
	})

	// Обработчик запроса на получение списка всех метрик
	r.GET("/", func(c *gin.Context) {
		//Собираем список метрик
		metricsList := ""

		memStorage.IterateMetrics(func(name string, value float64) {
			metricsList += fmt.Sprintf("Gauge: %s = %f\n", name, value)
		})
		memStorage.IterateCounters(func(name string, value int64) {
			metricsList += fmt.Sprintf("Counter: %s = %d\n", name, value)
		})

		// Если метрик нет, возвращаем сообщение
		if metricsList == "" {
			metricsList = "No metrics available."
		}

		// Отдаем HTML-страницу с метриками
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, "<html><body><h1>Metrics</h1><pre>%s</pre></body></html>", metricsList)
	})

	// Запускаем сервер на указанном адресе
	fmt.Printf("Server is running at http://%s\n", *serverAddr)
	if err := r.Run(*serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
