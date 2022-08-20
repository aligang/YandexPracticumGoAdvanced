package handler

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
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
		{
			name:  "FETCH ALL",
			input: input{path: "/", contentType: "text/plain"},
			expected: expected{
				code:        200,
				contentType: "text/html",
				payload: "gauge_example     1234\n" +
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
