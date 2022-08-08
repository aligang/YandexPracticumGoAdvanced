package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type expected struct {
	code        int
	contentType string
	value       string
}

type input struct {
	path        string
	contentType string
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
			expected: expected{code: 200, contentType: "text/plain", value: "1234"},
		},
		{
			name:     "CORRECT COUNTER",
			input:    input{path: "/value/counter/counter_example", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain", value: "12345"},
		},
		{
			name:     "Non EXISTING COUNTER",
			input:    input{path: "/value/counter/wrong_counter_example", contentType: "text/plain"},
			expected: expected{code: 404, contentType: "text/plain"},
		},
		{
			name:  "FETCH ALL",
			input: input{path: "/", contentType: "text/plain"},
			expected: expected{
				code:        200,
				contentType: "text/html",
				value: "gauge_example     1234\n" +
					"counter_example     12345\n",
			},
		},
	}

	strg := storage.Define(
		map[string]float64{
			"gauge_example": 1234,
		},
		map[string]int64{
			"counter_example": 12345,
		},
	)
	mux := New(strg)
	mux.Get("/value/{metricType}/{metricName}", mux.Fetch)
	mux.Get("/", mux.FetchAll)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, ts.URL+test.input.path, nil)
			require.NoError(t, err)
			request.Header.Add("Content-Type", test.input.contentType)
			res, err := http.DefaultClient.Do(request)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(res.Body)
			require.NoError(t, err)
			assert.Equal(t, test.expected.code, res.StatusCode)
			if res.StatusCode == http.StatusOK {
				assert.Equal(t, test.expected.contentType, res.Header.Get("Content-Type"))
				payload, err := io.ReadAll(res.Body)
				assert.NoError(t, err)
				assert.Equal(t, test.expected.value, string(payload))

			}
		})
	}
}
