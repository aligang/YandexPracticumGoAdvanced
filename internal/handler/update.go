package handler

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
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
	m := metric.Metrics{
		ID:    metricName,
		MType: metricType,
	}
	var updateErr error = nil
	if metricType == "gauge" {
		value, err := strconv.ParseFloat(metricValue, 64)
		m.Value = &value
		if err != nil {
			updateErr = err
			fmt.Println(err)
		}

	}
	if metricType == "counter" {
		value, err := strconv.ParseInt(metricValue, 10, 64)
		m.Delta = &value
		if err != nil {
			updateErr = err
			fmt.Println(err)
		}
	}
	if updateErr != nil {
		h.Storage.Update(m)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
