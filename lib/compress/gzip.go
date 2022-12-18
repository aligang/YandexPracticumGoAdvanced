package compress

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
)

// gzipWriter overloaded http.ResponseWriter that provides gzip compression/decompression
type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// GzipHandle middleware to provide gzip compression/decompression
func GzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		next.ServeHTTP(writer, r)
	})
}

//func AgentCompression(next func(r *http.Request) (*http.Response, error)) func(r *http.Request) (*http.Response, error) {
//	return func(r *http.Request) (*http.Response, error) {
//		gbuf := &bytes.Buffer{}
//		gz, err := gzip.NewWriterLevel(gbuf, gzip.BestSpeed)
//		if err != nil {
//			logging.Crit("Error During compressor creation")
//		}
//		res, err := io.ReadAll(r.Body)
//		if err != nil {
//			logging.Crit("Error During fetching data for compressiong")
//		}
//		_, err = gz.Write(res)
//		gz.Close()
//
//		if err != nil {
//			logging.Warn("Error During compression")
//		}
//
//		compressedRequest, err := http.NewRequest(r.Method, r.RequestURI, gbuf)
//		if err != nil {
//			logging.Warn("Error During creation of compressed request")
//		}
//		for r.Header.
//		compressedRequest.Header.
//		return next(compressedRequest)
//	}
//}
