package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
)

func (h APIHandler) UpdateWithJson(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		fmt.Println("Could not decode json")
		panic(err)
	}
	if m.MType != "gauge" && m.MType != "counter" {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}

	h.Storage.Update(m.ID, *m.Delta)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
