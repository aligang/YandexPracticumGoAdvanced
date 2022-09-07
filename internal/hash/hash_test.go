package hash

import (
	"crypto/hmac"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("HASHING of GAUGE", func(t *testing.T) {
		var refValue float64 = 4194304.000000
		refMetric := metric.Metrics{ID: "NextGC", MType: "gauge", Value: &refValue}
		refHash := "54b80493a21162a59b7362fee7445c12c6cab4fb912b8262aeadff13781e73a1"
		refKey := "/tmp/PHvif"
		hash, _ := CalculateHash(&refMetric, refKey)
		assert.Equal(t, true, hmac.Equal([]byte(refHash), []byte(hash)))
	})
}
