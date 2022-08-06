package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type expected struct {
	code        int
	contentType string
}

type input struct {
	path        string
	contentType string
}

func TestHandler(t *testing.T) {
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

		//{
		//	name:     "WRONG СOUNTER VALUE",
		//	input:    input{path: "/update/counter/PollCount/0.1", contentType: "text/plain"},
		//	expected: expected{code: 400, contentType: "text/plain"},
		//},
		{
			name:     "WRONG COUNTER VALUE",
			input:    input{path: "/update/counter/", contentType: "text/plain"},
			expected: expected{code: 404, contentType: "text/plain"},
		},
		{
			name:     "MAILFORMED СOUNTER VALUE",
			input:    input{path: "/update/counter/testCounter/none", contentType: "text/plain"},
			expected: expected{code: 400, contentType: "text/plain"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := ApiHandler{
				Storage: storage.New(),
			}
			request := httptest.NewRequest(http.MethodPost, test.input.path, nil)
			request.Header.Add("Content-Type", test.input.contentType)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(handler.ServeHTTP)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, test.expected.code, res.StatusCode)
			if res.StatusCode == http.StatusOK {
				assert.Equal(t, test.expected.contentType, res.Header.Get("Content-Type"))
			}
		})
	}

}
