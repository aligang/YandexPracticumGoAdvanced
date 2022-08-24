package handler

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h APIHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Processing incoming update to: %s\n", r.URL.String())
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")

	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		fmt.Printf("Got unsupported metric Type %s\n", metricType)
	}

	if !checkMetricValueFormat(metricType, metricValue) {
		http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
		fmt.Printf("Got unsupported metric format\n")
	}
	m := metric.Metrics{
		ID:    metricName,
		MType: metricType,
	}

	fmt.Printf("Parsing value for metric: %+v\n", m)
	if metricType == "gauge" {
		fmt.Println("Parsing gauge/float")
		value, _ := strconv.ParseFloat(metricValue, 64)
		m.Value = &value

	}
	if metricType == "counter" {
		fmt.Println("Parsing counter/int")
		value, _ := strconv.ParseInt(metricValue, 10, 64)
		m.Delta = &value
	}

	fmt.Printf("Updating storage with metric %+v\n", m)
	h.Storage.Update(m)

	fmt.Println("Prepare response")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Prepare response sent with status %d\n", http.StatusOK)
}
