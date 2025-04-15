package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/varik-08/go-metrics/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddMetricHandler(t *testing.T) {
	storage := internal.NewMemStorage()

	tests := []struct {
		url    string
		status int
		vars   map[string]string
	}{
		{
			url:    "/update/gauge/test/1",
			status: http.StatusOK,
			vars:   map[string]string{"name": "test", "type": "gauge", "value": "1"},
		},
		{
			url:    "/update/counter/test/1",
			status: http.StatusOK,
			vars:   map[string]string{"name": "test", "type": "counter", "value": "1"},
		},
		{
			url:    "/update/gauge/test/dsfds",
			status: http.StatusBadRequest,
			vars:   map[string]string{"name": "test", "type": "gauge", "value": "dsfds"},
		},
		{
			url:    "/update/unknown/testCounter/100'",
			status: http.StatusBadRequest,
			vars:   map[string]string{"name": "testCounter", "type": "unknown", "value": "100'"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, tt.url, nil)
			r = mux.SetURLVars(r, tt.vars)
			w := httptest.NewRecorder()

			addMetricHandler(w, r, storage)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}
