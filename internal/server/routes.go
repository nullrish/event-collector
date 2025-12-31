package server

import (
	"net/http"

	"github.com/nullrish/event-collector/internal/handler"
)

func registerRoutes(mux *http.ServeMux) {
	eventHandler := handler.NewEventHandler()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/events", eventHandler.Collect)
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("ok")); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
