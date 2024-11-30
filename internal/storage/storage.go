package storage

import "sync"

type MemStorage struct {
	metrics  map[string]float64
	counters map[string]int64
	mu       sync.RWMutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics:  make(map[string]float64),
		counters: make(map[string]int64),
	}
}

// функция-конструктор, которая создает новый экземпляр структуры.
//Она возвращает указатель на новый объект типа.

func (s *MemStorage) UpdateGauge(name string, value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.metrics[name] = value
}

func (s *MemStorage) UpdateCounter(name string, value int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counters[name] += value
}

//  методы позволяют извлекать данные о метриках, которые были ранее сохранены

func (s *MemStorage) GetGauge(name string) (float64, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, exists := s.metrics[name]
	return val, exists
}

func (s *MemStorage) GetCounter(name string) (int64, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, exist := s.counters[name]
	return val, exist
}

// IterateMetrics перебирает все метрики типа gauge.
func (s *MemStorage) IterateMetrics(fn func(name string, value float64)) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for name, value := range s.metrics {
		fn(name, value)
	}
}

// IterateCounters перебирает все счетчики.
func (s *MemStorage) IterateCounters(fn func(name string, value int64)) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for name, value := range s.counters {
		fn(name, value)
	}
}
