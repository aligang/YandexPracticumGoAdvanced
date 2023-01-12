package errors

import "errors"

var ErrInvalidMetricType = errors.New("invalid metric type")

var ErrInvalidHashValue = errors.New("invalid hash value")

var ErrRecordNotFound = errors.New("record not found")
