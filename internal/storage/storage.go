package storage

type MemStorage struct {
	metrics  map[string]float64
	counters map[string]int64
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
	s.metrics[name] = value
}

func (s *MemStorage) UpdateCounter(name string, value int64) {
	s.counters[name] += value
}

//  методы позволяют извлекать данные о метриках, которые были ранее сохранены

func (s *MemStorage) GetGauge(name string) (float64, bool) {
	val, exists := s.metrics[name]
	return val, exists
}

func (s *MemStorage) GetCounter(name string) (int64, bool) {
	val, exist := s.counters[name]
	return val, exist
}
