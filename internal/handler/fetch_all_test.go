package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchAll(t *testing.T) {
	tests := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name:  "FETCH ALL",
			input: input{path: "/", contentType: "text/plain"},
			expected: expected{
				code:        200,
				contentType: "text/html",
				payload: "{\"counter_example\":{\"id\":\"gauge_example\",\"type\":\"counter\",\"delta\":12345}," +
					"\"gauge_example\":{\"id\":\"gauge_example\",\"type\":\"gauge\",\"value\":1234}}",
			},
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
				assert.JSONEq(t, test.expected.payload, string(payload))
			}
		})
	}
}
