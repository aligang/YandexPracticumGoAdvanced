package encrypt

import (
	"bytes"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"io"
	"net/http"
)

func (k *EncryptionPlugin) EncryptWithPublicKey(next func(r *http.Request) (*http.Response, error)) func(r *http.Request) (*http.Response, error) {
	return func(r *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logging.Crit("Problem during reading request withing middleware")
		}
		encryptedBytes, err := EncryptOAEP(k.PublicKey, body)
		encryptedRequest, err := http.NewRequest(r.Method, r.RequestURI, bytes.NewBuffer(encryptedBytes))
		return next(encryptedRequest)
	}
}

func (k *EncryptionPlugin) DecryptWithPrivateKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if k.PrivateKey == nil {
			next.ServeHTTP(w, r)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logging.Crit("Problem during reading request withing middleware")
		}
		decryptedBytes, err := DecryptOAEP(k.PrivateKey, body)
		decryptedRequest, err := http.NewRequest(r.Method, r.RequestURI, bytes.NewBuffer(decryptedBytes))
		next.ServeHTTP(w, decryptedRequest)
	})
}
