package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func EncryptOAEP(public *rsa.PublicKey, rawMessage []byte) ([]byte, error) {
	label := []byte("OAEP Encrypted")
	random := rand.Reader
	hash := sha256.New()
	encryptedMessage := []byte{}

	messageLen := len(rawMessage)
	step := public.Size() - 2*hash.Size() - 2

	for start := 0; start < messageLen; start += step {
		finish := start + step
		if finish > messageLen {
			finish = messageLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash, random, public, rawMessage[start:finish], label)
		if err != nil {
			return encryptedMessage, err
		}
		encryptedMessage = append(encryptedMessage, encryptedBlockBytes...)
	}
	return encryptedMessage, nil
}

func DecryptOAEP(private *rsa.PrivateKey, message []byte) ([]byte, error) {
	label := []byte("OAEP Encrypted")
	random := rand.Reader
	hash := sha256.New()

	messageLen := len(message)
	step := private.PublicKey.Size()
	var decryptedBytes []byte

	for start := 0; start < messageLen; start += step {
		finish := start + step
		if finish > messageLen {
			finish = messageLen
		}

		decryptedBlockBytes, err := rsa.DecryptOAEP(hash, random, private, message[start:finish], label)
		if err != nil {
			return nil, err
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}

	return decryptedBytes, nil
}
