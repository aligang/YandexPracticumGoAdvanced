package handler

import (
	"fmt"
	. "github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h APIHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Processing incoming update to: %s\n", r.URL.String())
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")

	if !checkMetricType(&metricType) {
		http.Error(w, "Unsupported Metric Type", http.StatusNotImplemented)
		Logger.Warn().Msgf("Got unsupported metric Type %s\n", metricType)
		return
	}

	if !checkMetricValueFormat(metricType, metricValue) {
		http.Error(w, "Incorrect Metric Format", http.StatusBadRequest)
		Logger.Warn().Msgf("Got unsupported metric format\n")
		return
	}
	m := metric.Metrics{
		ID:    metricName,
		MType: metricType,
	}

	Logger.Debug().Msgf("Parsing value for metric: %+v\n", m)
	if metricType == "gauge" {
		Logger.Debug().Msgf("Parsing gauge/float")
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "Incorrect data format", http.StatusBadRequest)
			return
		}
		m.Value = &value

	}
	if metricType == "counter" {
		Logger.Debug().Msgf("Parsing counter/int")
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "Incorrect data format", http.StatusBadRequest)
			return
		}
		m.Delta = &value
	}
	Logger.Debug().Msgf("Value is metric: %+v\n", m)
	Logger.Debug().Msgf("Updating storage with metric %+v\n", m)
	h.Storage.Update(m)

	Logger.Debug().Msgf("Prepare response")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	Logger.Debug().Msgf("Prepare response sent with status %d\n", http.StatusOK)
}
