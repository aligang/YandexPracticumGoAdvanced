package handler

import (
	"bytes"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchWithJson(t *testing.T) {
	tests := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name: "CORRECT GAUGE",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"gauge_example\",\"type\":\"gauge\"}"},
			expected: expected{code: 200, contentType: "application/json",
				payload: "{\"id\":\"gauge_example\",\"type\":\"gauge\",\"value\":1234}"},
		},
		{
			name: "CORRECT COUNTER",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"counter_example\",\"type\":\"counter\"}"},
			expected: expected{code: 200, contentType: "application/json",
				payload: "{\"id\":\"counter_example\",\"type\":\"counter\",\"delta\":12345}"},
		},
	}

	var GaugeValue float64 = 1234
	var CounterDelta int64 = 12345
	strg := storage.New()
	strg.Load(
		map[string]metric.Metrics{
			"gauge_example":   {ID: "gauge_example", MType: "gauge", Value: &GaugeValue},
			"counter_example": {ID: "counter_example", MType: "counter", Delta: &CounterDelta},
		},
	)

	mux := New(strg)
	mux.Post("/value/", mux.FetchWithJson)
	ts := httptest.NewServer(mux)
	tc := ts.Client()
	defer ts.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodPost, ts.URL+test.input.path,
				bytes.NewBuffer([]byte(test.input.payload)))
			if err != nil {
				fmt.Println(err)
			}
			request.Header.Add("Content-Type", test.input.contentType)
			res, err := tc.Do(request)
			if err != nil {
				fmt.Println(err)
			}
			defer res.Body.Close()
			require.NoError(t, err)
			assert.Equal(t, test.expected.code, res.StatusCode)
			if res.StatusCode == http.StatusOK {
				assert.Equal(t, test.expected.contentType, res.Header.Get("Content-Type"))
				payload, err := io.ReadAll(res.Body)
				assert.NoError(t, err)
				assert.JSONEq(t, test.expected.payload, string(payload))

			}
		})
	}
}
