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

	err = h.BaseUpdate(m)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidMetricType):
			http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
			return
		case errors.Is(err, appErrors.ErrInvalidHashValue):
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
