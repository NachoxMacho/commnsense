package handler

import (
	"log/slog"
	"net/http"
)

func HTTPErrorHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "error", err, "path", r.URL.Path)
			w.Write([]byte("dum dum give me gum gum"))
			w.WriteHeader(500)
		}
	}
}
