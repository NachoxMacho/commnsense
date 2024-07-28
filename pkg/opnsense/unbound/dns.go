package unbound

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func GetDNSRecords() (records []DNSRecord, err error) {

	request, err := http.NewRequest("GET", "https://opnsense.robowens.dev/api/unbound/settings/searchHostOverride", nil)
	if err != nil {
		return nil, err
	}

	defaultTransport := http.DefaultTransport.(*http.Transport)

	// Create new Transport that ignores self-signed SSL
	customTransport := &http.Transport{
		Proxy:                 defaultTransport.Proxy,
		DialContext:           defaultTransport.DialContext,
		MaxIdleConns:          defaultTransport.MaxIdleConns,
		IdleConnTimeout:       defaultTransport.IdleConnTimeout,
		ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
		TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: customTransport}
	user := os.Getenv("OPNSENSE_USER")
	pass := os.Getenv("OPNSENSE_PASSWORD")
	request.SetBasicAuth(user, pass)
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := response.Body.Close()
		if closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	type OPNSenseData struct {
		Rows []apiDNSRecord `json:"rows"`
	}

	res := OPNSenseData{}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, fmt.Errorf("%w & json string %b", err, bodyBytes)
	}

	// fmt.Printf("Hello motherfucker")

	records = make([]DNSRecord, len(res.Rows))
	for i, l := range res.Rows {
		if records[i], err = convertAPIDNSRecord(l); err != nil {
			return nil, err
		}
	}

	return records, nil
}

func convertAPIDNSRecord(d apiDNSRecord) (DNSRecord, error) {

	record := DNSRecord{}
	record.Server = net.ParseIP(d.Server)
	if record.Server == nil {
		return DNSRecord{}, fmt.Errorf("invalid ip address in dns record %s", d.Server)
	}
	record.Description = d.Description
	record.Domain = d.Domain
	record.Enabled = d.Enabled == "1"
	record.Hostname = d.Hostname

	return record, nil
}
