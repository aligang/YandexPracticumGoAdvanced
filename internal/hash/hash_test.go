package hash

import (
	"crypto/hmac"
	"testing"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Run("HASHING of GAUGE", func(t *testing.T) {
		var refValue = 4194304.000000
		refMetric := metric.Metrics{ID: "NextGC", MType: "gauge", Value: &refValue}
		refHash := "54b80493a21162a59b7362fee7445c12c6cab4fb912b8262aeadff13781e73a1"
		refKey := "/tmp/PHvif"
		hash, _ := CalculateHash(&refMetric, refKey)
		assert.Equal(t, true, hmac.Equal([]byte(refHash), []byte(hash)))
	})
	t.Run("HASHING of COUNTER", func(t *testing.T) {
		var refDelta int64 = 3
		refMetric := metric.Metrics{ID: "PollCount", MType: "counter", Delta: &refDelta}
		refHash := "917f76e29b3cab53e3c9ea912567696501ec69bcbea708f89cf957a46b5b5012"
		refKey := "/tmp/e6qMTm"
		hash, _ := CalculateHash(&refMetric, refKey)
		assert.Equal(t, true, hmac.Equal([]byte(refHash), []byte(hash)))
	})
}
