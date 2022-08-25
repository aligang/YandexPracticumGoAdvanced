package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"io"
	"net/http"
)

func (h APIHandler) UpdateWithJSON(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read data", http.StatusUnsupportedMediaType)
	}
	err = json.Unmarshal(payload, &m)
	fmt.Println("Recieved JSON:")
	fmt.Println(string(payload))

	if err != nil {
		fmt.Println("Invalid JSON received")
		fmt.Println(err.Error())
		http.Error(w, "Invalid JSON received", http.StatusBadRequest)
	}
	if m.MType != "gauge" && m.MType != "counter" {
		fmt.Println("Invalid Metric Type")
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
	}
	h.Storage.Update(m)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
