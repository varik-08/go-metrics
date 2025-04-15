package main

import (
	"github.com/gorilla/mux"
	"github.com/varik-08/go-metrics/internal"
	"net/http"
	"strconv"
)

func addMetricHandler(w http.ResponseWriter, r *http.Request, memStorage internal.MemStorage) {
	vars := mux.Vars(r)
	nameMetric := vars["name"]
	typeMetric := vars["type"]

	switch typeMetric {
	case "gauge":
		value, err := strconv.ParseFloat(vars["value"], 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		memStorage.AddGauge(nameMetric, value)
	case "counter":
		value, err := strconv.Atoi(vars["value"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		memStorage.AddCounter(nameMetric, value)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
