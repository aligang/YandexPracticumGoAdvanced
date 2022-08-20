package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h APIHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.Storage.Dump()))
}

func (h APIHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}
	result, found := h.Storage.Get(metricType, metricName)
	w.Header().Set("Content-Type", "text/plain")
	if found {
		var reply string
		switch result.(type) {
		case int64:
			reply = fmt.Sprintf("%d", result.(int64))
		case float64:
			reply = strconv.FormatFloat(result.(float64), 'f', -1, 64)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(reply))
	} else {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", metricName), http.StatusNotFound)
	}

}
