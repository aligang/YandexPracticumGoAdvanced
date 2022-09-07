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
		refHash := "93d036cf9d13e15e0829338347901f16c158998cdc8156e0f08977625f30a462"
		refKey := "1"
		hash, _ := CalculateHash(&refMetric, refKey)
		assert.Equal(t, true, hmac.Equal([]byte(refHash), []byte(hash)))
	})
}
