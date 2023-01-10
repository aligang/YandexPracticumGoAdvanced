package handler

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

func (h *BaseHandler) BaseFetchAll() map[string]metric.Metrics {
	return h.Storage.Dump()
}
