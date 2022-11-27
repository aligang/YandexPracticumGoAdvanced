package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type expected struct {
	path        string
	code        int
	contentType string
	payload     string
}

type input struct {
	path        string
	contentType string
	payload     string
}

func TestFetch(t *testing.T) {
	tests := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name:     "CORRECT GAUGE",
			input:    input{path: "/value/gauge/gauge_example", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain", payload: "1234"},
		},
		{
			name:     "CORRECT COUNTER",
			input:    input{path: "/value/counter/counter_example", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain", payload: "12345"},
		},
		{
			name:     "Non EXISTING COUNTER",
			input:    input{path: "/value/counter/wrong_counter_example", contentType: "text/plain"},
			expected: expected{code: 404, contentType: "text/plain"},
		},
	}
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, ts.URL+test.input.path, nil)
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
				assert.Equal(t, test.expected.payload, string(payload))
			}
		})
	}
}
