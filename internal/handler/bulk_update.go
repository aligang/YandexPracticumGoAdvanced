package handler

import (
	"encoding/json"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
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
	logging.Debug("Received JSON: %s", string(payload))
	err = json.Unmarshal(payload, &metricSlice)
	if err != nil {
		logging.Warn("Invalid JSON received %s", err.Error())
		http.Error(w, "Invalid JSON received", http.StatusBadRequest)
		return
	}
	aggregatedMetrics := map[string]metric.Metrics{}
	for _, m := range metricSlice {
		if m.MType != "gauge" && m.MType != "counter" {
			logging.Warn("Invalid Metric Type")
			//http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
			continue
		}
		if h.Config.HashKey != "" {
			logging.Debug("Validating hash ...")
			if !hash.CheckHashInfo(&m, h.Config.HashKey) {
				logging.Warn("Invalid Hash")
				http.Error(w, "Invalid Hash", http.StatusBadRequest)
				continue
			} else {
				logging.Debug("Hash validation succeeded")
			}
		} else {
			logging.Debug("Skipping hash validation")
		}
		_, found := aggregatedMetrics[m.ID]
		if m.MType == "counter" && found {
			*aggregatedMetrics[m.ID].Delta += *m.Delta
		} else {
			aggregatedMetrics[m.ID] = m
		}
	}

	h.Storage.BulkUpdate(aggregatedMetrics)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
