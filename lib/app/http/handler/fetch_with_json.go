package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

// FetchWithJSON app API to download single metric with json-provided request
func (h APIHandler) FetchWithJSON(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics

	// DECRYPT_HERE

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		logging.Warn("Could not send byteData")
		http.Error(w, "Mailformed JSON", http.StatusBadRequest)
		return
	}
	if !checkMetricType(&m.MType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		return
	}
	result, found := h.Storage.Get(m.ID)
	if found {
		if h.Config.HashKey != "" {
			hash.AddHashInfo(&result, h.Config.HashKey)
		}
		j, err := json.Marshal(&result)
		if err != nil {
			logging.Warn("Could not encode Json")
			http.Error(w, "Mailformed JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		http.Error(w, fmt.Sprintf("Metric  with name=%s not found", m.ID), http.StatusNotFound)
	}

}
