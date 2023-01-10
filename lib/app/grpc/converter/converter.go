package converter

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/common"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

func ConvertMetric(in *common.Metric) metric.Metrics {
	delta := in.GetDelta()
	value := in.GetValue()
	hashValue := in.GetHash()

	m := metric.Metrics{
		ID:    in.GetID(),
		MType: in.GetMType(),
		Delta: &delta,
		Value: &value,
		Hash:  hashValue,
	}
	return m
}

func ConvertMetricEntity(e *metric.Metrics) *common.Metric {

	m := common.Metric{
		ID:           e.ID,
		MType:        e.MType,
		OptionalHash: &common.Metric_Hash{Hash: e.Hash},
	}
	if m.MType == "gauge" {
		m.OptionalValue = &common.Metric_Value{Value: *e.Value}
	} else if m.MType == "counter" {
		m.OptionalDelta = &common.Metric_Delta{Delta: *e.Delta}
	} else {
		return nil
	}
	return &m
}
