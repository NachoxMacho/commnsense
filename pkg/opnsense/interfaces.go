package opnsense

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Interface struct {
	Name string
}
type APIInterface struct {
	VlanTag interface{} `json:"vlan_tag,omitempty"`
	Config  struct {
		AdvDhcpConfigAdvanced         string `json:"adv_dhcp_config_advanced,omitempty"`
		AdvDhcpConfigFileOverride     string `json:"adv_dhcp_config_file_override,omitempty"`
		AdvDhcpConfigFileOverridePath string `json:"adv_dhcp_config_file_override_path,omitempty"`
		AdvDhcpOptionModifiers        string `json:"adv_dhcp_option_modifiers,omitempty"`
		AdvDhcpPtBackoffCutoff        string `json:"adv_dhcp_pt_backoff_cutoff,omitempty"`
		AdvDhcpPtInitialInterval      string `json:"adv_dhcp_pt_initial_interval,omitempty"`
		AdvDhcpPtReboot               string `json:"adv_dhcp_pt_reboot,omitempty"`
		AdvDhcpPtRetry                string `json:"adv_dhcp_pt_retry,omitempty"`
		AdvDhcpPtSelectTimeout        string `json:"adv_dhcp_pt_select_timeout,omitempty"`
		AdvDhcpPtTimeout              string `json:"adv_dhcp_pt_timeout,omitempty"`
		AdvDhcpPtValues               string `json:"adv_dhcp_pt_values,omitempty"`
		AdvDhcpRequestOptions         string `json:"adv_dhcp_request_options,omitempty"`
		AdvDhcpRequiredOptions        string `json:"adv_dhcp_required_options,omitempty"`
		AdvDhcpSendOptions            string `json:"adv_dhcp_send_options,omitempty"`
		AliasAddress                  string `json:"alias-address,omitempty"`
		AliasSubnet                   string `json:"alias-subnet,omitempty"`
		Blockbogons                   string `json:"blockbogons,omitempty"`
		Blockpriv                     string `json:"blockpriv,omitempty"`
		Descr                         string `json:"descr,omitempty"`
		Dhcphostname                  string `json:"dhcphostname,omitempty"`
		Dhcprejectfrom                string `json:"dhcprejectfrom,omitempty"`
		Enable                        string `json:"enable,omitempty"`
		Identifier                    string `json:"identifier,omitempty"`
		If                            string `json:"if,omitempty"`
		Ipaddr                        string `json:"ipaddr,omitempty"`
		Spoofmac                      string `json:"spoofmac,omitempty"`
		Lock                          string `json:"lock,omitempty"`
		Subnet                        string `json:"subnet,omitempty"`
		InternalDynamic               string `json:"internal_dynamic,omitempty"`
		Ipaddrv6                      string `json:"ipaddrv6,omitempty"`
		Subnetv6                      string `json:"subnetv6,omitempty"`
		Type                          string `json:"type,omitempty"`
		Virtual                       string `json:"virtual,omitempty"`
	} `json:"config,omitempty"`
	Statistics struct {
		HWOffloadCapabilities     string `json:"HW offload capabilities,omitempty"`
		AddressLength             string `json:"address length,omitempty"`
		BytesReceived             string `json:"bytes received,omitempty"`
		BytesTransmitted          string `json:"bytes transmitted,omitempty"`
		Collisions                string `json:"collisions,omitempty"`
		Datalen                   string `json:"datalen,omitempty"`
		Device                    string `json:"device,omitempty"`
		Driver                    string `json:"driver,omitempty"`
		Flags                     string `json:"flags,omitempty"`
		HeaderLength              string `json:"header length,omitempty"`
		Index                     string `json:"index,omitempty"`
		InputErrors               string `json:"input errors,omitempty"`
		InputQueueDrops           string `json:"input queue drops,omitempty"`
		LineRate                  string `json:"line rate,omitempty"`
		LinkState                 string `json:"link state,omitempty"`
		Metric                    string `json:"metric,omitempty"`
		Mtu                       string `json:"mtu,omitempty"`
		MulticastsReceived        string `json:"multicasts received,omitempty"`
		MulticastsTransmitted     string `json:"multicasts transmitted,omitempty"`
		OutputErrors              string `json:"output errors,omitempty"`
		PacketsForUnknownProtocol string `json:"packets for unknown protocol,omitempty"`
		PacketsReceived           string `json:"packets received,omitempty"`
		PacketsTransmitted        string `json:"packets transmitted,omitempty"`
		PromiscuousListeners      string `json:"promiscuous listeners,omitempty"`
		SendQueueDrops            string `json:"send queue drops,omitempty"`
		SendQueueLength           string `json:"send queue length,omitempty"`
		SendQueueMaxLength        string `json:"send queue max length,omitempty"`
		Type                      string `json:"type,omitempty"`
		UptimeAtAttachOrStatReset string `json:"uptime at attach or stat reset,omitempty"`
		Vhid                      string `json:"vhid,omitempty"`
	} `json:"statistics,omitempty"`
	Device            string        `json:"device,omitempty"`
	LinkType          string        `json:"link_type,omitempty"`
	Status            string        `json:"status,omitempty"`
	Description       string        `json:"description,omitempty"`
	Identifier        string        `json:"identifier,omitempty"`
	Mtu               string        `json:"mtu,omitempty"`
	MediaRaw          string        `json:"media_raw,omitempty"`
	Media             string        `json:"media,omitempty"`
	MacaddrHw         string        `json:"macaddr_hw,omitempty"`
	Macaddr           string        `json:"macaddr,omitempty"`
	Gateways          []interface{} `json:"gateways,omitempty"`
	Options           []interface{} `json:"options,omitempty"`
	Ipv6              []interface{} `json:"ipv6,omitempty"`
	Ipv4              []interface{} `json:"ipv4,omitempty"`
	IfctlSearchdomain []string      `json:"ifctl.searchdomain,omitempty"`
	IfctlRouter       []string      `json:"ifctl.router,omitempty"`
	IfctlNameserver   []string      `json:"ifctl.nameserver,omitempty"`
	Groups            []string      `json:"groups,omitempty"`
	Routes            []string      `json:"routes,omitempty"`
	Capabilities      []interface{} `json:"capabilities,omitempty"`
	Flags             []string      `json:"flags,omitempty"`
	SupportedMedia    []interface{} `json:"supported_media,omitempty"`
	Enabled           bool          `json:"enabled,omitempty"`
	IsPhysical        bool          `json:"is_physical,omitempty"`
}

func GetInterfaces(c Config) (interfaces []APIInterface, err error) {

	request, err := http.NewRequest("GET", c.BaseURL+"/api/interfaces/overview/export", nil)
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
		// TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	slog.Debug("Sending Request for interfaces", "Username", c.Authentication.Username, "| Password", c.Authentication.Password)

	client := &http.Client{Transport: customTransport}
	request.SetBasicAuth(c.Authentication.Username, c.Authentication.Password)
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

	res := []APIInterface{}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, fmt.Errorf("%w & json string %b", err, bodyBytes)
	}

	// // fmt.Printf("Hello motherfucker")
	//
	// records := make([]DNSRecord, len(res.Rows))
	// for i, l := range res.Rows {
	// 	if records[i], err = convertAPIDNSRecord(l); err != nil {
	// 		return nil, err
	// 	}
	// }
	//
	return res, nil

}
