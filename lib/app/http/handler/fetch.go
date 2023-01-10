package handler

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// Fetch app API to download single metric without payload-provided request
func (h HTTPHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		return
	}

	if metricType != "gauge" && metricType != "counter" {
		logging.Warn("Invalid Metric Type")
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		return
	}
	result, found := h.Storage.Get(metricName)

	if found {
		var reply string
		switch metricType {
		case "counter":
			reply = fmt.Sprintf("%d", *result.Delta)
		case "gauge":
			reply = strconv.FormatFloat(*result.Value, 'f', -1, 64)
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(reply))
		if err != nil {
			logging.Debug("Could not send byteData")
			http.Error(w, "Could send byteData", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", metricName), http.StatusNotFound)
	}

}
