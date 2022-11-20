package metric

//type Gauge float64
//type Counter int64
//
//type MemeStats struct {
//	Alloc         Gauge
//	BuckHashSys   Gauge
//	Frees         Gauge
//	GCCPUFraction Gauge
//	GCSys         Gauge
//	HeapAlloc     Gauge
//	HeapIdle      Gauge
//	HeapInuse     Gauge
//	HeapObjects   Gauge
//	HeapReleased  Gauge
//	HeapSys       Gauge
//	LastGC        Gauge
//	Lookups       Gauge
//	MCacheInuse   Gauge
//	MCacheSys     Gauge
//	MSpanInuse    Gauge
//	MSpanSys      Gauge
//	Mallocs       Gauge
//	NextGC        Gauge
//	NumForcedGC   Gauge
//	NumGC         Gauge
//	OtherSys      Gauge
//	PauseTotalNs  Gauge
//	StackInuse    Gauge
//	StackSys      Gauge
//	Sys           Gauge
//	TotalAlloc    Gauge
//}
//
//type OperStats struct {
//	PollCount   Counter
//	RandomValue Gauge
//}

type Stats struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	Hash  string   `json:"hash,omitempty"`
}

type MetricMap map[string]Metrics
