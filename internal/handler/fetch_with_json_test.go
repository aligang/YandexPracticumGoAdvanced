package handler

import (
	"bytes"
	"fmt"
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

	strg := storage.Define(
		map[string]float64{
			"gauge_example": 1234,
		},
		map[string]int64{
			"counter_example": 12345,
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
