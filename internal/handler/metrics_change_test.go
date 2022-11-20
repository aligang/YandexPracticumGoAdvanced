package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCounterIncrement(t *testing.T) {
	tests := []struct {
		name     string
		expected expected
		input    input
	}{
		{
			name: "UPDATE GAUGE 1",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"value\":323648,\"type\":\"gauge\"}"},
			expected: expected{code: 200, contentType: "text/plain",
				payload: ""},
		},
		{
			name: "CHECK GAUGE 1",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"type\":\"gauge\"}"},
			expected: expected{code: 200, contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"value\":323648,\"type\":\"gauge\"}"},
		},
		{
			name: "UPDATE GAUGE 2",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"value\":133,\"type\":\"gauge\"}"},
			expected: expected{code: 200, contentType: "text/plain",
				payload: ""},
		},
		{
			name: "CHECK GAUGE 2",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"type\":\"gauge\"}"},
			expected: expected{path: "/value/", code: 200, contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"value\":133,\"type\":\"gauge\"}"},
		},
		{
			name: "UPDATE COUNTER 1",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"delta\":10,\"type\":\"counter\"}"},
			expected: expected{code: 200, contentType: "text/plain",
				payload: ""},
		},
		{
			name: "CHECK COUNTER 1",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"type\":\"counter\"}"},
			expected: expected{code: 200, contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"delta\":10,\"type\":\"counter\"}"},
		},
		{
			name: "UPDATE COUNTER 2",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"delta\":10,\"type\":\"counter\"}"},
			expected: expected{code: 200, contentType: "text/plain",
				payload: ""},
		},
		{
			name: "CHECK COUNTER 2",
			input: input{path: "/value/", contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"type\":\"counter\"}"},
			expected: expected{code: 200, contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"delta\":20,\"type\":\"counter\"}"},
		},
	}

	strg := memory.New(nil)
	mux := New(strg, "", "Memory")
	mux.Post("/update/", mux.UpdateWithJSON)
	mux.Post("/value/", mux.FetchWithJSON)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request, err := http.NewRequest(http.MethodPost, ts.URL+test.input.path,
				bytes.NewBuffer([]byte(test.input.payload)))
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
			if len(test.expected.payload) > 0 {
				data, _ := io.ReadAll(res.Body)
				assert.JSONEq(t, test.expected.payload, string(data))
			}
		})
	}

}
