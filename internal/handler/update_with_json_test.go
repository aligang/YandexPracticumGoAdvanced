package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateWithJson(t *testing.T) {
	tests := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name: "CORRECT GAUGE",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"TotalAlloc\",\"value\":323648,\"type\":\"gauge\"}"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
		{
			name: "CORRECT COUNTER",
			input: input{path: "/update/", contentType: "application/json",
				payload: "{\"id\":\"PollCount\",\"delta\":10,\"type\":\"counter\"}"},
			expected: expected{code: 200, contentType: "text/plain"},
		},
	}

	strg := memory.New(nil)
	mux := New(strg, "", "Memory")
	mux.Post("/update/", mux.UpdateWithJSON)
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
		})
	}

}
