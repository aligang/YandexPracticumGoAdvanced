package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h APIHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	output, err := json.Marshal(h.Storage.Dump())
	if err != nil {
		fmt.Println("Problem During serialization of database")
	} else {
		w.Write(output)
	}
}

func (h APIHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}
	m := metric.Metrics{
		ID:    metricName,
		MType: metricType,
	}
	result, found := h.Storage.Get(m.ID)
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
		w.Write([]byte(reply))
	} else {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", metricName), http.StatusNotFound)
	}

}
