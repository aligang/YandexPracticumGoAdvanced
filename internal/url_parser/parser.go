package url_parser

import (
	"errors"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/url"
	"strings"
)

type parsedUrl struct {
	prefix string
	metric metric.Metric
}

func ParseUrl(url *url.URL) (metric.Metric, error) {
	_parsedUrl := parsedUrl{
		metric: metric.Metric{},
	}

	pathElem := strings.Split(url.Path, "/")
	errorMsg := fmt.Sprintf("Invalid Url Format: %s. ", url.Path)
	var errorDescription string
	if len(pathElem) != 5 {
		errorDescription = "Should be in \"/update/{{metricType}}/{{metricName}}/{{value}}\" format"
		return _parsedUrl.metric, errors.New(errorMsg + "\n" + errorDescription)
	} else {
		_parsedUrl.prefix = pathElem[1]
		_parsedUrl.metric.MetricType = pathElem[2]
		_parsedUrl.metric.MetricName = pathElem[3]
		_parsedUrl.metric.MetricValue = pathElem[4]
	}

	var parseError error = nil
	if !checkPrefix(_parsedUrl) {
		errorDescription = "Wrong prefix"
		parseError = errors.New(errorMsg + errorDescription)
	} else if !checkMetricType(_parsedUrl) {
		errorDescription = "Unsupported Metric Type"
		parseError = errors.New(errorMsg + errorDescription)
	} else if !checkMetricName(_parsedUrl) {
		errorDescription = "Unsupported Metric name"
		parseError = errors.New(errorMsg + errorDescription)
	} else if !checkMetricValueFormat(_parsedUrl) {
		errorDescription = "Mismatch between Metric Value and Metric Type"
		parseError = errors.New(errorMsg + errorDescription)
	}
	return _parsedUrl.metric, parseError
}
