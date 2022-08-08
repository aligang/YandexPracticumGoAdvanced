package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name:     "CORRECT GAUGE",
			input:    input{path: "/update/gauge/TotalAlloc/323648", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT GAUGE",
			input:    input{path: "/update/gauge/aaa/323648", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT COUNTER",
			input:    input{path: "/update/counter/PollCount/10", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT COUNTER",
			input:    input{path: "/update/counter/aaa/10", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT COUNTER",
			input:    input{path: "/update/counter/testCounter/100", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},

		{
			name:     "WRONG METRIC TYPE",
			input:    input{path: "/update/unknown/testCounter/100", contentType: "text/plain"},
			expected: expected{code: 501, contentType: "text/plain"},
		},
		{
			name:     "EMPTY COUNTER VALUE",
			input:    input{path: "/update/counter/", contentType: "text/plain"},
			expected: expected{code: 404, contentType: "text/plain"},
		},
		{
			name:     "MAILFORMED СOUNTER VALUE",
			input:    input{path: "/update/counter/testCounter/none", contentType: "text/plain"},
			expected: expected{code: 400, contentType: "text/plain"},
		},
	}

	strg := storage.New()
	mux := New(strg)
	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request, err := http.NewRequest(http.MethodPost, ts.URL+test.input.path, nil)
			require.NoError(t, err)
			request.Header.Add("Content-Type", test.input.contentType)
			res, err := http.DefaultClient.Do(request)
			require.NoError(t, err)
			assert.Equal(t, test.expected.code, res.StatusCode)
			if res.StatusCode == http.StatusOK {
				assert.Equal(t, test.expected.contentType, res.Header.Get("Content-Type"))
			}
		})
	}

}
