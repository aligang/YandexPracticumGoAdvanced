package hash

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("HASHING of GAUGE", func(t *testing.T) {
		var refValue float64 = 3877
		refMetric := metric.Metrics{ID: "BuckHashSys", MType: "gauge", Value: &refValue}
		refHash := "dd7677e2542cd3c1a170d66bad34a879f9027833f8894a76c729f58bcb55e5e8"
		refKey := "11111"
		hash, _ := CalculateHash(&refMetric, refKey)
		assert.Equal(t, refHash, hash)
	})
}
