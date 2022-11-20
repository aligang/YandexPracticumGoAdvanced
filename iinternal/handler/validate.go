package handler

import (
	"strconv"
)

func checkMetricValueFormat(metricType, metricValue string) bool {
	var err error = nil

	if metricType == "gauge" {
		_, err = strconv.ParseFloat(metricValue, 64)
	}
	if metricType == "counter" {
		_, err = strconv.ParseInt(metricValue, 10, 64)
	}
	if err != nil {
		return false
	}
	return true

}

func checkMetricType(metricType *string) bool {
	return *metricType == "gauge" || *metricType == "counter"
}
