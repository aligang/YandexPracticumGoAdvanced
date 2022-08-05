package url_parser

import (
	"strconv"
)

func checkPrefix(metricUrl parsedUrl) bool {
	return metricUrl.prefix == "update"
}

func checkMetricType(metricUrl parsedUrl) bool {
	return metricUrl.metric.MetricType == "gauge" || metricUrl.metric.MetricType == "counter"
}

func checkMetricName(metricUrl parsedUrl) bool {
	if metricUrl.metric.MetricName != "" {
		return true
	}
	return false
}

func checkMetricValueFormat(metricUrl parsedUrl) bool {
	_, ferr := strconv.ParseFloat(metricUrl.metric.MetricValue, 64)
	if ferr == nil {
		_, ierr := strconv.ParseInt(metricUrl.metric.MetricValue, 10, 64)
		if ierr == nil {
			if metricUrl.metric.MetricType == "counter" {
				return true
			}
		}
		if metricUrl.metric.MetricType == "gauge" {
			return true
		}
	}
	return false
}
