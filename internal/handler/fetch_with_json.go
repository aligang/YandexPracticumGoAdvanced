package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
)

func (h APIHandler) FetchWithJSON(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		fmt.Println("Could not decode json")
		http.Error(w, "Mailformed JSON", http.StatusBadRequest)
	}
	if !checkMetricType(&m.MType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}
	result, found := h.Storage.Get(m.ID)
	if found {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		j, err := json.Marshal(&result)
		if err != nil {
			fmt.Println("Could not encode Json")
			panic(err)
		}
		w.Write(j)
	} else {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", m.ID), http.StatusNotFound)
	}

}
