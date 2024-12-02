package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"metrics-tpl/internal/storage"
)

func TestMetricsHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/update/gauge/testMetric/123.45", nil)
	w := httptest.NewRecorder()

	// Создаем экземпляр MemStorage для передачи в обработчик
	storage := storage.NewMemStorage()
	handler := MetricsHandler(storage)

	// Вызываем обработчик
	handler(w, req)

	res := w.Result()
	defer res.Body.Close() // Закрываем тело ответа

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}

	// Дополнительные проверки, если необходимо
	if _, exists := storage.GetGauge("testMetric"); !exists {
		t.Error("Expected gauge 'testMetric' to exist in storage")
	}
}
