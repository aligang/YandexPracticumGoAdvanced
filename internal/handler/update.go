package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h ApiHandler) Update(w http.ResponseWriter, r *http.Request) {

	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")
	if metricType != "gauge" && metricType != "counter" {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}
	if !checkMetricValueFormat(metricType, metricValue) {
		http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
	}
	h.Storage.Update(metricType, metricName, metricValue)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
