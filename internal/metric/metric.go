package metric

import "runtime"

type MetricSet map[string]string

type Gauge float64
type Counter int64

func Metrics() MetricSet {
	return MetricSet{
		"Alloc":         "Gauge",
		"BuckHashSys":   "Gauge",
		"Frees":         "Gauge",
		"GCCPUFraction": "Gauge",
		"GCSys":         "Gauge",
		"HeapAlloc":     "Gauge",
		"HeapIdle":      "Gauge",
		"HeapInuse":     "Gauge",
		"HeapObjects":   "Gauge",
		"HeapReleased":  "Gauge",
		"HeapSys":       "Gauge",
		"LastGC":        "Gauge",
		"Lookups":       "Gauge",
		"MCacheInuse":   "Gauge",
		"MCacheSys":     "Gauge",
		"MSpanInuse":    "Gauge",
		"MSpanSys":      "Gauge",
		"Mallocs":       "Gauge",
		"NextGC":        "Gauge",
		"NumForcedGC":   "Gauge",
		"NumGC":         "Gauge",
		"OtherSys":      "Gauge",
		"PauseTotalNs":  "Gauge",
		"StackInuse":    "Gauge",
		"StackSys":      "Gauge",
		"Sys":           "Gauge",
		"TotalAlloc":    "Gauge",
	}
}

type OperStats struct {
	PollCount   Counter
	RandomValue Gauge
}

type Stats struct {
	MemStats  *runtime.MemStats
	OperStats *OperStats
}
