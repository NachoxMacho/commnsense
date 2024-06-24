package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/NachoxMacho/commnsense/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {

	if err := initAll(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/*", http.StripPrefix("/", http.FileServerFS(FS)))
	router.Get("/", handler.HTTPErrorHandler(handler.HandleHomeIndex))
	router.Get("/dropdown", handler.HTTPErrorHandler(handler.HandleDropDown))
	router.Post("/searchData", handler.HTTPErrorHandler(handler.HandleSearchData))

	slog.Info("server listening on", "port", os.Getenv("HTTP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), router))
}

// Specifically not just called init so we can load at a various point and do error handling
func initAll() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
