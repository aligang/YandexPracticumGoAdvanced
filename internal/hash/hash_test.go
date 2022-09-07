package hash

import (
	"crypto/hmac"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("HASHING of GAUGE", func(t *testing.T) {
		var refValue float64 = 3877
		refMetric := metric.Metrics{ID: "BuckHashSys", MType: "gauge", Value: &refValue}
		refHash := "dbb85a3637412267a14d69f51eb8e0dab2f918379994702a7b5ba381128a2d70"
		refKey := "1"
		hash, _ := CalculateHash(&refMetric, refKey)
		assert.Equal(t, true, hmac.Equal([]byte(refHash), []byte(hash)))
	})
}
