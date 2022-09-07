package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
)

func CalculateHash(m *metric.Metrics, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	var hashingMaterial string

	if m.MType == "gauge" {
		hashingMaterial = fmt.Sprintf("%s:counter:%d", m.ID, m.Delta)
	} else if m.MType == "counter" {
		hashingMaterial = fmt.Sprintf("%s:gauge:%d", m.ID, m.Value)
	}

	h.Write([]byte(hashingMaterial))
	dst := h.Sum(nil)
	return hex.EncodeToString(dst)
}

func AddHashInfo(m *metric.Metrics, key string) {
	m.Hash = CalculateHash(m, key)
}

func CheckHashInfo(m *metric.Metrics, key string) bool {
	return m.Hash == CalculateHash(m, key)
}
