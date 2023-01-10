package errors

import "errors"

//var InvMetric error = &InvalidMetricType{"invalid metric type"}

//type InvalidMetricType struct {
//	msg string
//}
//
//func (e *InvalidMetricType) Error() string {
//	return e.msg
//}

var InvalidMetricType error = errors.New("invalid metric type")

var InvalidHashValue error = errors.New("invalid hash value")

var RecordNotFound error = errors.New("record not found")
