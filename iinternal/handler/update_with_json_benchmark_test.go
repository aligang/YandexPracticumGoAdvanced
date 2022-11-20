package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/storage/memory"
)

func BenchmarkUpdateWithJson(b *testing.B) {

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
			name: "CORRECT GAUGE",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"value\":323648,\"type\":\"gauge\"}",
			},
		},

		{
			name: "CORRECT COUNTER",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"delta\":10,\"type\":\"counter\"}",
			},
		},
		{
			name: "CORRECT GAUGE",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"value\":323648,\"type\":\"gauge\"}",
			},
		},
		{
			name: "CORRECT COUNTER",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"delta\":10,\"type\":\"counter\"}",
			},
		},
	}

	var requests []*http.Request
	for _, test := range tests {
		request, _ := http.NewRequest(http.MethodPost, ts.URL+test.input.path, bytes.NewBuffer([]byte(test.input.payload)))
		request.Header.Add("Content-Type", test.input.contentType)
		requests = append(requests, request)
	}

	b.StartTimer()
	for _, request := range requests {
		b.StartTimer()
		res, _ := tc.Do(request)
		b.StopTimer()
		defer res.Body.Close()
	}

}
