package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/go-chi/chi/v5"
)

// Update server API to upload single metric without payload-provided request
func (h APIHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Processing incoming update to: %s\n", r.URL.String())
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")

	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		logging.Warn("Got unsupported metric Type %s\n", metricType)
		return
	}

	if !checkMetricValueFormat(metricType, metricValue) {
		http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
		logging.Warn("Got unsupported metric format\n")
		return
	}
	m := metric.Metrics{
		ID:    metricName,
		MType: metricType,
	}

	logging.Debug("Parsing value for metric: %+v\n", m)
	if metricType == "gauge" {
		logging.Debug("Parsing gauge/float")
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "Incorrect data format", http.StatusBadRequest)
			return
		}
		m.Value = &value

	}
	if metricType == "counter" {
		logging.Debug("Parsing counter/int")
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "Incorrect data format", http.StatusBadRequest)
			return
		}
		m.Delta = &value
	}
	logging.Debug("Value is metric: %+v\n", m)
	logging.Debug("Updating storage with metric %+v\n", m)
	err := h.Storage.Update(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	logging.Debug("Prepare response")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	logging.Debug("Prepare response sent with status %d\n", http.StatusOK)
}
