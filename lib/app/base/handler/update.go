package handler

import (
	"fmt"
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

func (h *BaseHandler) BaseUpdate(m metric.Metrics) error {

	if m.MType != "gauge" && m.MType != "counter" {
		logging.Warn("Invalid Metric Type")
		return appErrors.ErrInvalidMetricType
	}
	if h.Config.HashKey != "" {
		logging.Debug("Validating hash ...")
		if !hash.CheckHashInfo(&m, h.Config.HashKey) {
			logging.Warn("Invalid Hash")
			return appErrors.ErrInvalidHashValue
		} else {
			logging.Debug("Hash validation succeeded")
		}
	} else {
		logging.Debug("Skipping hash validation")
	}
	err := h.Storage.Update(m)
	if err != nil {
		return fmt.Errorf(`database communication error: %w`, err)
	}
	return nil
}
