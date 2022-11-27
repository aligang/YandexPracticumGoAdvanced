package handler

import (
	"net/http"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
)

// Ping server API to check connectivity ot repository
func (h APIHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if h.Config.StorageType == "Memory" {
		logging.Warn("Ping Handler is not supported for current storage type")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if h.Config.StorageType == "Database" {
		err := h.Storage.IsAlive()
		if err != nil {
			logging.Warn("DB Storage connection id dead")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	logging.Debug("DB Storage is alive")
}
