package handler

import (
	appErrors "github.com/aligang/YandexPracticumGoAdvanced/lib/app/base/errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

func (h *BaseHandler) BaseFetch(m metric.Metrics) (*metric.Metrics, error) {

	if m.MType != "gauge" && m.MType != "counter" {
		logging.Warn("Invalid Metric Type")
		return nil, appErrors.ErrInvalidMetricType
	}
	result, found := h.Storage.Get(m.ID)
	if found && h.Config.HashKey != "" {
		hash.AddHashInfo(&result, h.Config.HashKey)
	}
	if found {
		return &result, nil
	}
	return nil, appErrors.ErrRecordNotFound

}
