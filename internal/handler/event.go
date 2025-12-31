// Package handler..
package handler

import "net/http"

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (h *EventHandler) Collect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO
	if _, err := w.Write([]byte("LEFT TO DO")); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
