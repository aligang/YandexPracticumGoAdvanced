package converter

import (
	"errors"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"strconv"
)

func ConvertPlainMetric(metricID string, metricType string, metricValue string) (*metric.Metrics, error) {
	if !checkMetricValueFormat(metricType, metricValue) {
		logging.Warn("Got unsupported metric format\n")
		return nil, errors.New("Got unsupported metric format\n")
	}
	m := &metric.Metrics{
		ID:    metricID,
		MType: metricType,
	}

	logging.Debug("Parsing value for metric: %+v\n", m)
	if metricType == "gauge" {
		logging.Debug("Parsing gauge/float")
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return nil, errors.New("Got unsupported metric format\n")
		}
		m.Value = &value

	}
	if metricType == "counter" {
		logging.Debug("Parsing counter/int")
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return nil, errors.New("incorrect data format")
		}
		m.Delta = &value
	}
	return m, nil
}

func ConvertMetricEntityToPlain(m *metric.Metrics) string {
	var reply string
	switch m.MType {
	case "counter":
		reply = fmt.Sprintf("%d", *m.Delta)
	case "gauge":
		reply = strconv.FormatFloat(*m.Value, 'f', -1, 64)
	}
	return reply
}
