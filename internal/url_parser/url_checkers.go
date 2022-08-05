package url_parser

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"strconv"
)

func checkPrefix(metricUrl parsedUrl) bool {
	return metricUrl.prefix == "update"
}

func checkMetricType(metricUrl parsedUrl) bool {
	return metricUrl.metric.MetricType == "gauge" || metricUrl.metric.MetricType == "counter"
}

func checkMetricName(metricUrl parsedUrl) bool {
	metricTypes := metric.GetMetricTypes()
	_, found := metricTypes[metricUrl.metric.MetricName]
	if found {
		return true
	}
	return false
}

func checkMetricValueType(metricUrl parsedUrl) bool {
	metricTypes := metric.GetMetricTypes()
	metricType, found := metricTypes[metricUrl.metric.MetricName]
	if found {
		if metricType == metricUrl.metric.MetricType {
			return true
		}
	}
	return false
}

func checkMetricValueFormat(metricUrl parsedUrl) bool {
	metricTypes := metric.GetMetricTypes()
	metricType, _ := metricTypes[metricUrl.metric.MetricName]
	_, ferr := strconv.ParseFloat(metricUrl.metric.MetricValue, 64)
	if ferr == nil {
		_, ierr := strconv.ParseInt(metricUrl.metric.MetricValue, 10, 64)
		if ierr == nil {
			if metricType == "counter" {
				return true
			}
		}
		if metricType == "gauge" {
			return true
		}
	}
	return false
}
