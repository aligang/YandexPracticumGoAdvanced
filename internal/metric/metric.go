package metric

import (
	"reflect"
	"strings"
)

type Metric struct {
	MetricType  string
	MetricName  string
	MetricValue string
}

type Gauge float64
type Counter int64

type MemeStats struct {
	Alloc         Gauge
	BuckHashSy    Gauge
	Frees         Gauge
	GCCPUFraction Gauge
	GCSys         Gauge
	HeapAlloc     Gauge
	HeapIdle      Gauge
	HeapInuse     Gauge
	HeapObjects   Gauge
	HeapReleased  Gauge
	HeapSys       Gauge
	LastGC        Gauge
	Lookups       Gauge
	MCacheInuse   Gauge
	MCacheSys     Gauge
	MSpanInuse    Gauge
	MSpanSys      Gauge
	Mallocs       Gauge
	NextGC        Gauge
	NumForcedGC   Gauge
	NumGC         Gauge
	OtherSys      Gauge
	PauseTotalNs  Gauge
	StackInuse    Gauge
	StackSys      Gauge
	Sys           Gauge
	TotalAlloc    Gauge
}

type OperStats struct {
	PollCount   Counter
	RandomValue Gauge
}

type Stats struct {
	MemeStats
	OperStats
}

func GetMetricTypes() map[string]string {
	metricTypes := map[string]string{}

	s := reflect.ValueOf(Stats{})
	for i := 0; i < s.NumField(); i++ {
		e := s.Field(i)
		for j := 0; j < e.NumField(); j++ {
			metricValue := e.Field(j)
			metricName := e.Type().Field(j).Name
			metricType := strings.ToLower(metricValue.Type().Name())
			metricTypes[metricName] = metricType
		}
	}
	return metricTypes
}
