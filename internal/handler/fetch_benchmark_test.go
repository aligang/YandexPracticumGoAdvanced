package handler

import (
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage/memory"
)

func BenchmarkFetch(b *testing.B) {
	b.StopTimer()

	var GaugeValue float64 = 1234
	var CounterDelta int64 = 12345

	strg := memory.New(nil)
	strg.Load(
		map[string]metric.Metrics{
			"gauge_example":   {ID: "gauge_example", MType: "gauge", Value: &GaugeValue},
			"counter_example": {ID: "gauge_example", MType: "counter", Delta: &CounterDelta},
		},
	)
	mux := New(strg, "", "Memory")
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
	mux.Get("/", mux.FetchAll)
	ts := httptest.NewServer(mux)
	tc := ts.Client()
	defer ts.Close()

	tests := []struct {
		name  string
		input input
	}{
		{
			name:  "CORRECT GAUGE",
			input: input{path: "/value/gauge/gauge_example", contentType: "text/plain"},
		},
		{
			name:  "CORRECT COUNTER",
			input: input{path: "/value/counter/counter_example", contentType: "text/plain"},
		},
	}

	var requests []*http.Request
	for _, test := range tests {
		request, _ := http.NewRequest(http.MethodGet, ts.URL+test.input.path, nil)
		request.Header.Add("Content-Type", test.input.contentType)
		requests = append(requests, request)
	}

	for _, request := range requests {
		b.StartTimer()
		res, _ := tc.Do(request)
		b.StopTimer()
		defer res.Body.Close()
	}

}
