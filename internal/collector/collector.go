package collector

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"math/rand"
	"runtime"
	"time"
)

func CollectMetrics(m *metric.Stats) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	for {
		runtime.ReadMemStats(m.MemStats)
		m.OperStats.PollCount += 1
		m.OperStats.RandomValue = metric.Gauge(r.Float64())

		//fmt.Println("data polled")
		time.Sleep(2 * time.Second)
	}
}
