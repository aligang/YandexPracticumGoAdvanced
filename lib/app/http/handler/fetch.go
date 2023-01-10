package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/go-chi/chi/v5"
)

// FetchAll app API to download all metrics
func (h APIHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	output, err := json.Marshal(h.Storage.Dump())
	if err != nil {
		logging.Warn("Problem During serialization of database")
		http.Error(w, "Problem During serialization of database", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

// Fetch app API to download single metric without payload-provided request
func (h APIHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		return
	}
	result, found := h.Storage.Get(metricName)
	w.Header().Set("Content-Type", "text/plain")
	if found {
		var reply string
		switch metricType {
		case "counter":
			reply = fmt.Sprintf("%d", *result.Delta)
		case "gauge":
			reply = strconv.FormatFloat(*result.Value, 'f', -1, 64)
		}
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
