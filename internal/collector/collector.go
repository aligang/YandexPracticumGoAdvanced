package collector

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"math/rand"
	"runtime"
	"time"
)

func CollectOperStats(m *metric.Stats, r *rand.Rand) {
	m.Counter["PollCount"] += 1
	m.Gauge["RandomValue"] = r.Float64()
}

func CollectMemStats(m *metric.Stats) {
	memstats := &runtime.MemStats{}
	runtime.ReadMemStats(memstats)

	m.Gauge["Alloc"] = float64(memstats.Alloc)
	m.Gauge["BuckHashSys"] = float64(memstats.BuckHashSys)
	m.Gauge["Frees"] = float64(memstats.Frees)
	m.Gauge["GCCPUFraction"] = float64(memstats.GCCPUFraction)
	m.Gauge["GCSys"] = float64(memstats.GCSys)
	m.Gauge["HeapAlloc"] = float64(memstats.HeapAlloc)
	m.Gauge["HeapIdle"] = float64(memstats.HeapIdle)
	m.Gauge["HeapInuse"] = float64(memstats.HeapInuse)
	m.Gauge["HeapObjects"] = float64(memstats.HeapObjects)
	m.Gauge["HeapReleased"] = float64(memstats.HeapReleased)
	m.Gauge["HeapSys"] = float64(memstats.HeapSys)
	m.Gauge["LastGC"] = float64(memstats.LastGC)
	m.Gauge["Lookups"] = float64(memstats.Lookups)
	m.Gauge["MCacheInuse"] = float64(memstats.MCacheInuse)
	m.Gauge["MCacheSys"] = float64(memstats.MCacheSys)
	m.Gauge["MSpanInuse"] = float64(memstats.MSpanInuse)
	m.Gauge["MSpanSys"] = float64(memstats.MSpanSys)
	m.Gauge["Mallocs"] = float64(memstats.Mallocs)
	m.Gauge["NextGC"] = float64(memstats.NextGC)
	m.Gauge["NumForcedGC"] = float64(memstats.NumForcedGC)
	m.Gauge["NumGC"] = float64(memstats.NumGC)
	m.Gauge["OtherSys"] = float64(memstats.OtherSys)
	m.Gauge["PauseTotalNs"] = float64(memstats.PauseTotalNs)
	m.Gauge["StackInuse"] = float64(memstats.StackInuse)
	m.Gauge["StackSys"] = float64(memstats.StackSys)
	m.Gauge["Sys"] = float64(memstats.Sys)
	m.Gauge["TotalAlloc"] = float64(memstats.TotalAlloc)
}

func CollectMetrics(m *metric.Stats) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	ticker := time.NewTicker(2 * time.Second)
	for {
		CollectMemStats(m)
		CollectOperStats(m, r)

		fmt.Println("data polled")

		<-ticker.C
	}
}
