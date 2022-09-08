package handler

import (
	"fmt"
	"net/http"
)

func (h APIHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if h.Config.StorageType == "Memory" {
		fmt.Printf("Ping Handler is not supported for current storage type")
		http.Error(w, "", http.StatusInternalServerError)
	}

	if h.Config.StorageType == "Database" {
		err := h.Storage.IsAlive()
		if err != nil {
			fmt.Printf("DB Storage connection id dead")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Printf("DB Storage is alive")
}
