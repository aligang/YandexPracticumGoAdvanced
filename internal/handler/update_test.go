package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name:     "CORRECT GAUGE 1",
			input:    input{path: "/update/gauge/TotalAlloc/323648", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT GAUGE 2",
			input:    input{path: "/update/gauge/aaa/323648", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT COUNTER 1",
			input:    input{path: "/update/counter/PollCount/10", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT COUNTER 2",
			input:    input{path: "/update/counter/vvv/10", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name:     "CORRECT COUNTER 3",
			input:    input{path: "/update/counter/testCounter/100", contentType: "text/plain"},
			expected: expected{code: 200, contentType: "text/plain"},
		},

		{
			name:     "WRONG METRIC TYPE",
			input:    input{path: "/update/unknown/testCounter/100", contentType: "text/plain"},
			expected: expected{code: 501, contentType: "text/plain"},
		},
		{
			name:     "EMPTY COUNTER VALUE ",
			input:    input{path: "/update/counter/", contentType: "text/plain"},
			expected: expected{code: 404, contentType: "text/plain"},
		},
		{
			name:     "MAILFORMED СOUNTER VALUE",
			input:    input{path: "/update/counter/testCounter/none", contentType: "text/plain"},
			expected: expected{code: 400, contentType: "text/plain"},
		},
	}

	strg := memory.New(nil)
	mux := New(strg, "", "Memory")
	mux.Post("/update/{metricType}/{metricName}/{metricValue}", mux.Update)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodPost, ts.URL+test.input.path, nil)
			require.NoError(t, err)
			request.Header.Add("Content-Type", test.input.contentType)
			res, err := http.DefaultClient.Do(request)
			if err != nil {
				fmt.Println(err)
			}
			defer res.Body.Close()
			require.NoError(t, err)
			assert.Equal(t, test.expected.code, res.StatusCode)
			if res.StatusCode == http.StatusOK {
				assert.Equal(t, test.expected.contentType, res.Header.Get("Content-Type"))
			}
		})
	}

}
