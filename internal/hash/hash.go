package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	. "github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
)

func CalculateHash(m *metric.Metrics, key string) (string, error) {
	var hashingMaterial string
	var err error

	h := hmac.New(sha256.New, []byte(key))

	switch m.MType {
	case "counter":
		hashingMaterial = fmt.Sprintf("%s:%s:%d", m.ID, m.MType, *m.Delta)
		Logger.Debug().Msgf("Hashing material is: %s\n", hashingMaterial)
	case "gauge":
		hashingMaterial = fmt.Sprintf("%s:%s:%f", m.ID, m.MType, *m.Value)
	default:
		return "", err
	}
	Logger.Debug().Msgf("Hashing material is : %s\n", hashingMaterial)
	_, err = h.Write([]byte(hashingMaterial))
	if err != nil {
		Logger.Warn().Msgf("Could not Calculate hash for: %v\n", *m)
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
		Logger.Warn().Msgf("Could not add hash to: %v\n", *m)
		return
	}

	m.Hash = hash
	Logger.Debug().Msgf("Applying hash for: %v\n", *m)
}

func CheckHashInfo(m *metric.Metrics, key string) bool {
	hash, err := CalculateHash(m, key)
	if err != nil {
		Logger.Warn().Msgf("Could not calculate hash for: %v\n", *m)
		return false
	}
	Logger.Warn().Msgf("Checking provided hash for: %v\n", *m)
	res := hmac.Equal([]byte(m.Hash), []byte(hash))
	if res {
		Logger.Debug().Msgf("Hash for: %+v is valid\n", *m)
	} else {
		Logger.Warn().Msgf("Hash for: %+v is invalid\n", *m)
	}
	return res
}
