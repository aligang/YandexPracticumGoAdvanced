package handler

import (
	"encoding/json"
	"errors"
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"io"
	"net/http"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

// BulkUpdate app API to upload multiple metrics
func (h HTTPHandler) BulkUpdate(w http.ResponseWriter, r *http.Request) {
	var metricSlice []metric.Metrics

	// DECRYPT_HERE

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read data", http.StatusUnsupportedMediaType)
		logging.Warn("Could not read data %s", err.Error())
		return
	}
	logging.Debug("Received JSON: %s", string(payload))
	err = json.Unmarshal(payload, &metricSlice)
	if err != nil {
		logging.Warn("Invalid JSON received %s", err.Error())
		http.Error(w, "Invalid JSON received", http.StatusBadRequest)
		return
	}
	err = h.BaseBulkUpdate(metricSlice)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.InvalidMetricType):
			http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
			return
		case errors.Is(err, appErrors.InvalidHashValue):
			http.Error(w, "Invalid Hash", http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
