package collector

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"math/rand"
	"runtime"
	"time"
)

func CollectOperStats(os *metric.Stats, r *rand.Rand) {
	os.PollCount += 1
	os.RandomValue = metric.Gauge(r.Float64())
}

func CollectMemStats(m *metric.Stats) {
	memstats := &runtime.MemStats{}
	runtime.ReadMemStats(memstats)

	m.Alloc = metric.Gauge(memstats.Alloc)
	m.BuckHashSy = metric.Gauge(memstats.BuckHashSys)
	m.Frees = metric.Gauge(memstats.Frees)
	m.GCCPUFraction = metric.Gauge(memstats.GCCPUFraction)
	m.GCSys = metric.Gauge(memstats.GCSys)
	m.HeapAlloc = metric.Gauge(memstats.HeapAlloc)
	m.HeapIdle = metric.Gauge(memstats.HeapIdle)
	m.HeapInuse = metric.Gauge(memstats.HeapInuse)
	m.HeapObjects = metric.Gauge(memstats.HeapObjects)
	m.HeapReleased = metric.Gauge(memstats.HeapReleased)
	m.HeapSys = metric.Gauge(memstats.HeapSys)
	m.LastGC = metric.Gauge(memstats.LastGC)
	m.Lookups = metric.Gauge(memstats.Lookups)
	m.MCacheInuse = metric.Gauge(memstats.MCacheInuse)
	m.MCacheSys = metric.Gauge(memstats.MCacheSys)
	m.MSpanInuse = metric.Gauge(memstats.MSpanInuse)
	m.MSpanSys = metric.Gauge(memstats.MSpanSys)
	m.Mallocs = metric.Gauge(memstats.Mallocs)
	m.NextGC = metric.Gauge(memstats.NextGC)
	m.NumForcedGC = metric.Gauge(memstats.NumForcedGC)
	m.NumGC = metric.Gauge(memstats.NumGC)
	m.OtherSys = metric.Gauge(memstats.OtherSys)
	m.PauseTotalNs = metric.Gauge(memstats.PauseTotalNs)
	m.StackInuse = metric.Gauge(memstats.StackInuse)
	m.StackSys = metric.Gauge(memstats.StackSys)
	m.Sys = metric.Gauge(memstats.Sys)
	m.TotalAlloc = metric.Gauge(memstats.TotalAlloc)
}

func CollectMetrics(m *metric.Stats) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	for {
		CollectMemStats(m)
		CollectOperStats(m, r)

		fmt.Println("data polled")
		time.Sleep(2 * time.Second)
	}
}
