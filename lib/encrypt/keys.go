package encrypt

import (
	"crypto/rsa"
	"crypto/x509"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"io"
	"os"
)

func ReadPublicKey(filePath string) *rsa.PublicKey {
	if filePath == "" {
		return nil
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		logging.Crit("Could not open file for reading")
		panic(err.Error())
	}
	keyBytes, err := io.ReadAll(file)
	if err != nil {
		logging.Warn("Could not read public key from file")
		panic(err.Error())
	}
	//block, _ := pem.Decode(keyBytes)
	//pubclicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	publicKey, err := x509.ParsePKCS1PublicKey(keyBytes)
	if err != nil {
		logging.Crit("Could not parse public key")
		panic(err.Error())
	}
	return publicKey
}

func ReadPrivateKey(filePath string) *rsa.PrivateKey {
	if filePath == "" {
		return nil
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		logging.Warn("Could not open file for reading")
		panic(err.Error())
	}
	keyBytes, err := io.ReadAll(file)
	if err != nil {
		logging.Warn("Could not read public key from file")
		panic(err.Error())
	}
	//block, _ := pem.Decode(keyBytes)
	//pubclicKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err != nil {
		logging.Crit("Could not parse private key")
		return nil
	}
	return privateKey
}
