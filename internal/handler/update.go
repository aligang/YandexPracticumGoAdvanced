package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h APIHandler) Update(w http.ResponseWriter, r *http.Request) {

	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")

	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}

	if !checkMetricValueFormat(metricType, metricValue) {
		http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
	}

	if metricType == "gauge" {
		value, err := strconv.ParseFloat(metricValue, 64)
		if err == nil {
			h.Storage.UpdateGauge(metricName, value)
		} else {
			fmt.Println(err)
		}
	}
	if metricType == "counter" {
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err == nil {
			h.Storage.UpdateCounter(metricName, value)
		} else {
			fmt.Println(err)
		}
	}

	//h.Storage.Update(metricName, metricValue)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
