package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"io"
	"net/http"
)

func (h APIHandler) BulkUpdate(w http.ResponseWriter, r *http.Request) {
	var metricSlice []metric.Metrics

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read data", http.StatusUnsupportedMediaType)
		return
	}
	fmt.Println("Recieved JSON:")
	fmt.Println(string(payload))
	err = json.Unmarshal(payload, &metricSlice)
	if err != nil {
		fmt.Println("Invalid JSON received")
		fmt.Println(err.Error())
		http.Error(w, "Invalid JSON received", http.StatusBadRequest)
		return
	}
	for _, m := range metricSlice {
		if m.MType != "gauge" && m.MType != "counter" {
			fmt.Println("Invalid Metric Type")
			http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
			return
		}
		if h.Config.HashKey != "" {
			fmt.Println("Validating hash ...")
			if !hash.CheckHashInfo(&m, h.Config.HashKey) {
				fmt.Println("Invalid Hash")
				http.Error(w, "Invalid Hash", http.StatusBadRequest)
				return
			} else {
				fmt.Println("Hash validation succeeded")
			}
		} else {
			fmt.Println("Skipping hash validation")
		}

		h.Storage.Update(m)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
