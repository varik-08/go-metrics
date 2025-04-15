package internal

import "testing"
import "github.com/stretchr/testify/assert"

func TestMemStorage_AddCounter(t *testing.T) {
	storage := NewMemStorage()

	tests := []struct {
		name  string
		value int
	}{
		{
			name:  "test 1",
			value: 1,
		},
		{
			name:  "test 100",
			value: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage.AddCounter(tt.name, tt.value)

			assert.Equal(t, tt.value, storage.GetCounter(tt.name))
		})
	}
}

func TestMemStorage_AddGauge(t *testing.T) {
	storage := NewMemStorage()

	tests := []struct {
		name  string
		value float64
	}{
		{
			name:  "test 1",
			value: 1,
		},
		{
			name:  "test 100.6",
			value: 100.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage.AddGauge(tt.name, tt.value)

			assert.Equal(t, tt.value, storage.GetGauge(tt.name))
		})
	}
}
