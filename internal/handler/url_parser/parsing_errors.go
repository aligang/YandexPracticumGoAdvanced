package url_parser

import "errors"

var WrongUrlFormat = errors.New("WrongUrl Format. Should be in \"/update/{{metricType}}/{{metricName}}/{{value}} format")
var EmptyValueError = errors.New("ValueIsEmpty")
var MailformedValueError = errors.New("Wrong ID format")
var UnsupportedMetricType = errors.New("Wrong Metric type")
