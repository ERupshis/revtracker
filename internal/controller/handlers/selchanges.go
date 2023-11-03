package handlers

import (
	"fmt"
	"net/http"

	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
)

func SelectChanges(storage storage.BaseStorage, log logger.BaseLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respBody := []byte("TODO")

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Length", fmt.Sprintf("%d", len(respBody)))
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write(respBody); err != nil {
			log.Info("[controller:handlers:Balance] failed to write orders data in response body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
