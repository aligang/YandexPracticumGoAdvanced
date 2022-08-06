package url_parser

import (
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
	if len(pathElem) != 5 {
		return _parsedUrl.metric, EmptyValueError
	} else {
		_parsedUrl.prefix = pathElem[1]
		_parsedUrl.metric.MetricType = pathElem[2]
		_parsedUrl.metric.MetricName = pathElem[3]
		_parsedUrl.metric.MetricValue = pathElem[4]
	}

	var parseError error = nil
	if !checkPrefix(_parsedUrl) {
		parseError = WrongUrlFormat
	} else if !checkMetricType(_parsedUrl) {
		parseError = WrongUrlFormat
	} else if !CheckValueIsNotEmpty(_parsedUrl) {
		parseError = EmptyValueError
	} else if !checkMetricValueFormat(_parsedUrl) {
		parseError = MailformedValueError
	}
	return _parsedUrl.metric, parseError
}
