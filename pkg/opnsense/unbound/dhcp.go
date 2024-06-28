package unbound

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

func GetDHCPLeases() ([]Lease, error) {

	request, err := http.NewRequest("GET", "https://opnsense.robowens.dev/api/dhcpv4/leases/searchLease", nil)
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
	request.SetBasicAuth("8ewFwkMxUq2FTckEGiDjUzLcgqaMbFDmwommJGaJaV1VlXYQi8X/Ibgir++DAJ8HbhxognU7J0Pmvn9F", "2fIiZpYRwHH2TcrAQAFXCzgWXrD3CLW9M3XvB621xz5gKUHT6xU4v6Rf/+gyGBhL0AZ0KkWPIagHCdFz")
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	type OPNSenseData struct {
		Rows []apiLease `json:"rows"`
	}

	res := OPNSenseData{}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, err
	}

	leases := make([]Lease, len(res.Rows))
	for i, l := range res.Rows {
		if leases[i], err = convertAPILease(l); err != nil {
			return nil, err
		}
	}

	return leases, nil
}

func convertAPILease(l apiLease) (Lease, error) {

	newLease := Lease{}
	var err error

	newLease.IP = net.ParseIP(l.Address)
	if newLease.IP == nil {
		return Lease{}, fmt.Errorf("ip address not valid %s", l.Address)
	}

	newLease.MAC, err = net.ParseMAC(l.Mac)
	if err != nil {
		return Lease{}, err
	}

	newLease.Description = l.Descr
	newLease.Interface = l.If
	newLease.Ends = l.Ends
	newLease.DeviceModel = l.Man
	newLease.Hostname = l.Hostname
	newLease.Type = l.Type
	newLease.State = l.State
	newLease.Status = l.Status

	return newLease, nil
}
