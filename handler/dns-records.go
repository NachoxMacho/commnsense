package handler

import (
	"net/http"

	"github.com/NachoxMacho/commnsense/pkg/k8s"
	"github.com/NachoxMacho/commnsense/pkg/opnsense/unbound"
	"github.com/NachoxMacho/commnsense/view/dns"
)

func HandleDNSRecords(w http.ResponseWriter, r *http.Request) error {
	dnsRecords, err := unbound.GetDNSRecords()
	if err != nil {
		return err
	}

	ingresses, err := k8s.GetIngresses()
	if err != nil {
		return err
	}

	return dns.Index(dnsRecords, ingresses).Render(r.Context(), w)
}
