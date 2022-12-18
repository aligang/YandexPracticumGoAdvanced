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
		if err != nil {
			logging.Crit("Problem during data encryption")
		}
		encryptedRequest, err := http.NewRequest(r.Method, r.RequestURI, bytes.NewBuffer(encryptedBytes))
		if err != nil {
			logging.Crit("Problem during adding of encrypted payload to request")
		}
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
		if err != nil {
			logging.Crit("Problem during data decryption")
		}
		decryptedRequest, err := http.NewRequest(r.Method, r.RequestURI, bytes.NewBuffer(decryptedBytes))
		if err != nil {
			logging.Crit("Problem during adding of decrypted payload to request")
		}
		next.ServeHTTP(w, decryptedRequest)
	})
}
