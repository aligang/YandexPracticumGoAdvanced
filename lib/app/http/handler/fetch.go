package handler

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/http/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Fetch app API to download single metric without payload-provided request
func (h HTTPHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	dummyValue := "0"
	m, err := converter.ConvertPlainMetric(metricName, metricType, dummyValue)
	if err != nil {
		http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
		return
	}

	result, err := h.BaseFetch(*m)
	if err != nil {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", m.ID), http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	reply := converter.ConvertMetricEntityToPlain(result)
	_, err = w.Write([]byte(reply))
	if err != nil {
		logging.Debug("Could not send byteData")
		http.Error(w, "Could send byteData", http.StatusInternalServerError)
		return
	}
}
