package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
)

func CalculateHash(m *metric.Metrics, key string) (string, error) {
	var hashingMaterial string
	var err error

	h := hmac.New(sha256.New, []byte(key))

	if m.MType == "gauge" {
		hashingMaterial = fmt.Sprintf("%s:counter:%d", m.ID, m.Delta)
	} else if m.MType == "counter" {
		hashingMaterial = fmt.Sprintf("%s:gauge:%d", m.ID, m.Value)
	}

	_, err = h.Write([]byte(hashingMaterial))
	if err != nil {
		fmt.Printf("Could not Calculate hash for: %v\n", *m)
		return "", err
	}

	dst := h.Sum(nil)
	hash := hex.EncodeToString(dst)
	fmt.Printf("Calculated hash for: %v = %s\n", *m, hash)
	return hash, nil
}

func AddHashInfo(m *metric.Metrics, key string) {
	hash, err := CalculateHash(m, key)
	if err != nil {
		fmt.Printf("Could not add hash to: %v\n", *m)
		return
	}

	m.Hash = hash
	fmt.Printf("Applying hash for: %v\n", *m)
}

func CheckHashInfo(m *metric.Metrics, key string) bool {
	hash, err := CalculateHash(m, key)
	if err != nil {
		fmt.Printf("Could not calculate hash for: %v\n", *m)
		return false
	}
	fmt.Printf("Checking provided hash for: %v\n", *m)
	res := m.Hash == hash
	if res {
		fmt.Printf("Hash for: %+v is valid\n", *m)
	} else {
		fmt.Printf("Hash for: %+v is invalid\n", *m)
	}
	return res
}

func test() {

}
