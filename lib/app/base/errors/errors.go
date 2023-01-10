package errors

import "errors"

var InvalidMetricType = errors.New("invalid metric type")

var InvalidHashValue = errors.New("invalid hash value")

var RecordNotFound = errors.New("record not found")
