package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/hash"
	. "github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
)

func (h APIHandler) FetchWithJSON(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		Logger.Warn().Msg("Could not send byteData")
		http.Error(w, "Mailformed JSON", http.StatusBadRequest)
		return
	}
	if !checkMetricType(&m.MType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		return
	}
	result, found := h.Storage.Get(m.ID)
	if found {
		if len(result.Hash) == 0 && h.Config.HashKey != "" {
			hash.AddHashInfo(&result, h.Config.HashKey)
		}
		j, err := json.Marshal(&result)
		if err != nil {
			Logger.Warn().Msg("Could not encode Json")
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
