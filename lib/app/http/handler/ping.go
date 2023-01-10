package handler

import (
	"net/http"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
)

// Ping app API to check connectivity ot repository
func (h HTTPHandler) Ping(w http.ResponseWriter, r *http.Request) {
	err := h.BasePing()
	if err != nil {
		logging.Warn(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	logging.Debug("DB Storage is alive")
}
