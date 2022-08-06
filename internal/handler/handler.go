package handler

import (
	"errors"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/handler/url_parser"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"net/http"
)

type ApiHandler struct {
	Storage *storage.Storage
}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported Http Method", http.StatusBadRequest)
	}

	metric, err := url_parser.ParseUrl(r.URL)
	if err != nil {
		switch {
		case errors.Is(err, url_parser.EmptyValueError):
			http.Error(w, fmt.Sprintln(err), http.StatusNotFound)
		case errors.Is(err, url_parser.MailformedValueError):
			http.Error(w, fmt.Sprintln(err), http.StatusBadRequest)
		case errors.Is(err, url_parser.WrongUrlFormat):
			http.Error(w, fmt.Sprintln(err), http.StatusNotImplemented)
		case errors.Is(err, url_parser.UnsupportedMetricType):
			http.Error(w, fmt.Sprintln(err), http.StatusNotImplemented)
		}
	}

	h.Storage.Update(&metric)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Println("2")
}
