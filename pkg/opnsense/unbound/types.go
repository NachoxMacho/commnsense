package unbound

import "net"

type Lease struct {
	Description string
	Ends        string
	Interface   string
	State       string
	Status      string
	DeviceModel string
	Hostname    string
	Type        string
	IP          net.IP
	MAC         net.HardwareAddr
}

type apiLease struct {
	Address  string `json:"address,omitempty"`
	Descr    string `json:"descr,omitempty"`
	Ends     string `json:"ends,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	If       string `json:"if,omitempty"`
	IfDescr  string `json:"if_descr,omitempty"`
	Mac      string `json:"mac,omitempty"`
	Man      string `json:"man,omitempty"`
	Starts   string `json:"starts,omitempty"`
	State    string `json:"state,omitempty"`
	Status   string `json:"status,omitempty"`
	Type     string `json:"type,omitempty"`
}

type DNSRecord struct {
	Description string
	Domain      string
	Hostname    string
	Server      net.IP
	Enabled     bool
}

type apiDNSRecord struct {
	Description string `json:"description,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Enabled     string `json:"enabled,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Mx          string `json:"mx,omitempty"`
	Mxprio      string `json:"mxprio,omitempty"`
	Rr          string `json:"rr,omitempty"`
	Server      string `json:"server,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}
