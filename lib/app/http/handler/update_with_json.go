package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

// UpdateWithJSON app API to upload single metric with json-formated request
func (h HTTPHandler) UpdateWithJSON(w http.ResponseWriter, r *http.Request) {
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
	//COMMON PART
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
	//COMMON PART

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
