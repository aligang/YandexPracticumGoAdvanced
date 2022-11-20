package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/storage/memory"
)

func BenchmarkFetchWithJson(b *testing.B) {
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
			name: "CORRECT GAUGE",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"gauge_example\",\"type\":\"gauge\"}"},
		},

		{
			name: "CORRECT COUNTER",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"counter_example\",\"type\":\"counter\"}",
			},
		},
	}

	var requests []*http.Request
	for _, test := range tests {
		request, _ := http.NewRequest(http.MethodPost, ts.URL+test.input.path, bytes.NewBuffer([]byte(test.input.payload)))
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
