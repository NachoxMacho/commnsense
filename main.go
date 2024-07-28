package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/NachoxMacho/commnsense/handler"
	"github.com/NachoxMacho/commnsense/pkg/opnsense"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {

	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env, continuing with existing environment variables")
	}

	opnsenseConfig, err := opnsense.NewConfig(
		opnsense.WithURL("https://opnsense.robowens.dev"),
		opnsense.WithAuthentication(os.Getenv("OPNSENSE_USERNAME"), os.Getenv("OPNSENSE_PASSWORD")),
	)
	if err != nil {
		slog.Error("Invalid Opnsense Configuration", slog.Attr{ Key: "Error", Value: slog.AnyValue(err) })
	}

	handlerConfig := handler.Config{
		OpnSense: opnsenseConfig,
	}

	log.Println("OPNSENSE URL:", opnsenseConfig.BaseURL)

	if err := initAll(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/*", http.StripPrefix("/", http.FileServerFS(FS)))
	router.Get("/", handler.HTTPErrorHandler(handler.HandleNewHomeIndex(handlerConfig)))
	router.Get("/dropdown", handler.HTTPErrorHandler(handler.HandleDropDown(handlerConfig)))
	router.Post("/searchData", handler.HTTPErrorHandler(handler.HandleSearchData(handlerConfig)))
	router.Get("/dns", handler.HTTPErrorHandler(handler.HandleDNSRecords))

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
