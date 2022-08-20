package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"io"
	"net/http"
)

func (h APIHandler) UpdateWithJson(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics
	//decoder := json.NewDecoder(r.Body)
	//err := decoder.Decode(&m)
	payload, err := io.ReadAll(r.Body)
	err = json.Unmarshal(payload, &m)
	if err != nil {
		fmt.Println("Could not decode json")
		panic(err)
	}
	if m.MType != "gauge" && m.MType != "counter" {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}

	switch m.MType {
	case "gauge":
		h.Storage.UpdateGauge(m.ID, *m.Value)
		fmt.Println(*m.Value)
	case "counter":
		h.Storage.UpdateCounter(m.ID, *m.Delta)
		fmt.Println(*m.Delta)
	default:
		fmt.Println("miss")
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
