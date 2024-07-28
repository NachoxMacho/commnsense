package handler

import (
	"net/http"

	"github.com/NachoxMacho/commnsense/pkg/opnsense"
)

type Config struct {
	OpnSense opnsense.Config
}

type httpHandler func(http.ResponseWriter, *http.Request) error
