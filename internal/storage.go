package internal

type MemStorage struct {
	gauges   map[string]float64
	counters map[string]int
}

func NewMemStorage() MemStorage {
	return MemStorage{
		gauges:   make(map[string]float64),
		counters: make(map[string]int),
	}
}

func (m *MemStorage) GetGauge(name string) float64 {
	return m.gauges[name]
}

func (m *MemStorage) GetCounter(name string) int {
	return m.counters[name]
}

func (m *MemStorage) AddGauge(name string, value float64) {
	m.gauges[name] = value
}

func (m *MemStorage) AddCounter(name string, value int) {
	m.counters[name] += value
}
