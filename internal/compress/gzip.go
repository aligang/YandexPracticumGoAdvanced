package compress

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
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
			} else {
				r.Body = gz
				fmt.Println("Decompression was applied")
			}
		} else {
			fmt.Println("No Compression in request header")
		}

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			writer = w
			fmt.Println("No Response compressiong will be provided")
		} else {
			gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			} else {
				defer gz.Close()
				writer = gzipWriter{ResponseWriter: w, Writer: gz}
				fmt.Println("Response will be compressed")
				writer.Header().Set("Content-Encoding", "gzip")
			}
		}
		next(writer, r)
	}
}
