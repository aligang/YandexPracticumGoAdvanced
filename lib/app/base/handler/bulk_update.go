package handler

import (
	"fmt"
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

func (h *BaseHandler) BaseBulkUpdate(metrics []metric.Metrics) error {
	aggregatedMetrics := map[string]metric.Metrics{}
	for _, m := range metrics {
		if m.MType != "gauge" && m.MType != "counter" {
			logging.Warn("Invalid Metric Type")
			return appErrors.ErrInvalidMetricType
		}
		if h.Config.HashKey != "" {
			logging.Debug("Validating hash ...")
			if !hash.CheckHashInfo(&m, h.Config.HashKey) {
				logging.Warn("Invalid Hash")
				return appErrors.ErrInvalidHashValue
			}
			logging.Debug("Hash validation succeeded")
		} else {
			logging.Debug("Skipping hash validation")
		}
		_, found := aggregatedMetrics[m.ID]
		if m.MType == "counter" && found {
			*aggregatedMetrics[m.ID].Delta += *m.Delta
		} else {
			aggregatedMetrics[m.ID] = m
		}
	}
	err := h.Storage.BulkUpdate(aggregatedMetrics)
	if err != nil {
		return fmt.Errorf(`database communication error: %w`, err)
	}
	return nil
}
