package handler

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/url_parser"
	"net/http"
)

type ApiHandler struct {
	Storage *storage.Storage
}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported Http Method", http.StatusBadRequest)
	}

	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, "Unsupported Content Type", http.StatusBadRequest)
	}

	metric, err := url_parser.ParseUrl(r.URL)
	if err != nil {
		http.Error(w, fmt.Sprintln(err), http.StatusBadRequest)
	}

	h.Storage.Update(&metric)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Println("2")
}
