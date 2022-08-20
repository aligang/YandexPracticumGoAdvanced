package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
)

func (h APIHandler) FetchWithJson(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		fmt.Println("Could not decode json")
		panic(err)
	}
	if !checkMetricType(&m.MType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}
	result, found := h.Storage.Get(m.MType, m.ID)
	w.Header().Set("Content-Type", "application/json")
	if found {
		var reply metric.Metrics
		switch m.MType {
		case "gauge":
			value := result.(float64)
			reply = metric.Metrics{ID: m.ID, MType: m.MType, Value: &value}
		case "counter":
			value := result.(int64)
			reply = metric.Metrics{ID: m.ID, MType: m.MType, Delta: &value}
		}

		w.WriteHeader(http.StatusOK)
		j, err := json.Marshal(&reply)
		if err != nil {
			fmt.Println("Could not encode Json")
			panic(err)
		}
		w.Write(j)
	} else {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", m.ID), http.StatusNotFound)
	}

}
