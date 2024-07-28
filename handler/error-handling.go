package handler

import (
	"log/slog"
	"net/http"
)

func HTTPErrorHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "error", err, "path", r.URL.Path)
			_, err := w.Write([]byte("dum dum give me gum gum"))
			if err != nil {
				slog.Error("failed to write response", slog.Any("error", err))
			}
			w.WriteHeader(500)
		}
	}
}
