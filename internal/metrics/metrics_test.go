package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sega-strn/metrics-tpl/internal/storage" // Импортируем пакет storage
)

func TestMetricsHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/update/gauge/testMetric/123.45", nil)
	w := httptest.NewRecorder()

	// Создаем экземпляр MemStorage для передачи в обработчик
	storage := storage.NewMemStorage() // Убедитесь, что эта функция определена
	handler := MetricsHandler(storage)

	// Вызываем обработчик
	handler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}

	// Дополнительные проверки, если необходимо
	if _, exists := storage.GetGauge("testMetric"); !exists {
		t.Error("Expected gauge 'testMetric' to exist in storage")
	}
}
