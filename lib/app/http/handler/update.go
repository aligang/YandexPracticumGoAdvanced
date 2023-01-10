package handler

import (
	"errors"
	"fmt"
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/http/converter"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Update app API to upload single metric without payload-provided request
func (h HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Processing incoming update to: %s\n", r.URL.String())
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")
	//
	//if !converter.checkMetricValueFormat(metricType, metricValue) {
	//	http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
	//	logging.Warn("Got unsupported metric format\n")
	//	return
	//}
	//m := metric.Metrics{
	//	ID:    metricName,
	//	MType: metricType,
	//}
	//
	//logging.Debug("Parsing value for metric: %+v\n", m)
	//if metricType == "gauge" {
	//	logging.Debug("Parsing gauge/float")
	//	value, err := strconv.ParseFloat(metricValue, 64)
	//	if err != nil {
	//		http.Error(w, "Incorrect data format", http.StatusBadRequest)
	//		return
	//	}
	//	m.Value = &value
	//
	//}
	//if metricType == "counter" {
	//	logging.Debug("Parsing counter/int")
	//	value, err := strconv.ParseInt(metricValue, 10, 64)
	//	if err != nil {
	//		http.Error(w, "Incorrect data format", http.StatusBadRequest)
	//		return
	//	}
	//	m.Delta = &value
	//}
	m, err := converter.ConvertPlainMetric(metricName, metricType, metricValue)
	if err != nil {
		http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
		return
	}
	err = h.BaseUpdate(*m)
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
	logging.Debug("Prepare response")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	logging.Debug("Prepare response sent with status %d\n", http.StatusOK)
}
