package compress

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/logging"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GzipHandle(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var writer http.ResponseWriter

		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, "Incorrect Encoding", http.StatusBadRequest)
				return
			} else {
				r.Body = gz
				logging.Debug("Decompression was applied")
			}
		} else {
			logging.Debug("No Compression in request header")
		}

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			writer = w
			logging.Debug("No Response compression will be provided")
		} else {
			gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			writer = gzipWriter{ResponseWriter: w, Writer: gz}
			logging.Debug("Response will be compressed")
			writer.Header().Set("Content-Encoding", "gzip")
			defer gz.Close()
		}
		next(writer, r)
	}
}
