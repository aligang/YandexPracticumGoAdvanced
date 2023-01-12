package handler

import (
	"errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
)

func (h *BaseHandler) BasePing() error {
	if h.Config.StorageType == "Memory" {
		logging.Warn("Ping Handler is not supported for current storage type")
		return errors.New("ping handler is not supported for current storage type")
	}

	if h.Config.StorageType == "Database" {
		err := h.Storage.IsAlive()
		if err != nil {
			logging.Warn("DB Storage connection id dead")
			return errors.New("DB Storage connection id dead")
		}
	}
	logging.Debug("DB Storage is alive")
	return nil
}
