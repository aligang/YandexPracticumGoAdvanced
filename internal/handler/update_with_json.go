package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
)

func (h APIHandler) UpdateWithJSON(w http.ResponseWriter, r *http.Request) {
	var m metric.Metrics

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read data", http.StatusUnsupportedMediaType)
		return
	}
	logging.Debug("Received JSON:")
	logging.Debug(string(payload))
	err = json.Unmarshal(payload, &m)
	if err != nil {
		logging.Warn("Invalid JSON received %s", err.Error())
		http.Error(w, "Invalid JSON received", http.StatusBadRequest)
		return
	}
	if m.MType != "gauge" && m.MType != "counter" {
		logging.Warn("Invalid Metric Type")
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		return
	}
	if h.Config.HashKey != "" {
		logging.Debug("Validating hash ...")
		if !hash.CheckHashInfo(&m, h.Config.HashKey) {
			logging.Warn("Invalid Hash")
			http.Error(w, "Invalid Hash", http.StatusBadRequest)
			return
		} else {
			logging.Debug("Hash validation succeeded")
		}
	} else {
		logging.Debug("Skipping hash validation")
	}

	err = h.Storage.Update(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
