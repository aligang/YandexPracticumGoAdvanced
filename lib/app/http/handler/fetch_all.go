package handler

import (
	"encoding/json"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"net/http"
)

// FetchAll app API to download all metrics
func (h HTTPHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	output, err := json.Marshal(h.BaseFetchAll())
	if err != nil {
		logging.Warn("Problem During serialization of database")
		http.Error(w, "Problem During serialization of database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
