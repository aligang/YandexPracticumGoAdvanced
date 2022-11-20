package handler

import (
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/storage/memory"
)

func BenchmarkUpdate(b *testing.B) {
	b.StopTimer()

	strg := memory.New(nil)
	mux := New(strg, "", "Memory")
	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	ts := httptest.NewServer(mux)
	tc := ts.Client()
	defer func() {
		b.StopTimer()
		defer ts.Close()
	}()

	tests := []struct {
		name  string
		input input
	}{
		{
			name:  "CORRECT GAUGE 1",
			input: input{path: "/update/gauge/TotalAlloc/323648", contentType: "text/plain"},
		},
		{
			name:  "CORRECT GAUGE 2",
			input: input{path: "/update/gauge/aaa/323648", contentType: "text/plain"},
		},
		{
			name:  "CORRECT COUNTER 1",
			input: input{path: "/update/counter/PollCount/10", contentType: "text/plain"},
		},
		{
			name:  "CORRECT COUNTER 2",
			input: input{path: "/update/counter/vvv/10", contentType: "text/plain"},
		},
	}

	var requests []*http.Request
	for _, test := range tests {
		request, _ := http.NewRequest(http.MethodPost, ts.URL+test.input.path, nil)
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
