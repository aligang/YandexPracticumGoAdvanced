package handler

import (
	. "github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"net/http"
)

func (h APIHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if h.Config.StorageType == "Memory" {
		Logger.Warn().Msg("Ping Handler is not supported for current storage type")
		http.Error(w, "", http.StatusInternalServerError)
	}

	if h.Config.StorageType == "Database" {
		err := h.Storage.IsAlive()
		if err != nil {
			Logger.Warn().Msg("DB Storage connection id dead")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	Logger.Debug().Msg("DB Storage is alive")
}
