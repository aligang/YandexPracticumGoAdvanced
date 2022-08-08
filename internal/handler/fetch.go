package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h ApiHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.Storage.Dump()))
}

func (h ApiHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	if metricType != "gauge" && metricType != "counter" {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}
	result, found := h.Storage.Get(metricType, metricName)
	w.Header().Set("Content-Type", "text/plain")
	if found {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	} else {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", metricName), http.StatusNotFound)
	}

}
