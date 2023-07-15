package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type TrojanConfig struct {
	Type      string           `json:"type"`
	Tag       string           `json:"tag"`
	Server    string           `json:"server"`
	Port      int              `json:"server_port"`
	Password  string           `json:"password"`
	Network   string           `json:"network"`
	TLS       TLSOptions       `json:"tls"`
	Multiplex MultiplexOptions `json:"multiplex"`
	Transport TransportOptions `json:"transport"`
}

type VMessConfig struct {
	Type      string           `json:"type"`
	Tag       string           `json:"tag"`
	Users     []User           `json:"users"`
	TLS       TLSOptions       `json:"tls"`
	Transport TransportOptions `json:"transport"`
}

type VLESSConfig struct {
	Type      string           `json:"type"`
	Tag       string           `json:"tag"`
	ID        string           `json:"id"`
	Server    string           `json:"server"`
	Port      int              `json:"server_port"`
	TLS       TLSOptions       `json:"tls"`
	Transport TransportOptions `json:"transport"`
}

type TransportOptions interface{}
type User struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	AlterID int    `json:"alterId"`
}
type TLSOptions struct {
	Enabled         bool     `json:"enabled"`
	DisableSNI      bool     `json:"disable_sni"`
	ServerName      string   `json:"server_name"`
	Insecure        bool     `json:"insecure"`
	ALPN            []string `json:"alpn"`
	MinVersion      string   `json:"min_version"`
	MaxVersion      string   `json:"max_version"`
	CipherSuites    []string `json:"cipher_suites"`
	Certificate     string   `json:"certificate"`
	CertificatePath string   `json:"certificate_path"`
	ECH             ECHOptions
	UTLS            UTLSOptions
	Reality         RealityOptions
}

type ECHOptions struct {
	Enabled                     bool   `json:"enabled"`
	PQSignatureSchemesEnabled   bool   `json:"pq_signature_schemes_enabled"`
	DynamicRecordSizingDisabled bool   `json:"dynamic_record_sizing_disabled"`
	Config                      string `json:"config"`
}

type UTLSOptions struct {
	Enabled     bool   `json:"enabled"`
	Fingerprint string `json:"fingerprint"`
}

type RealityOptions struct {
	Enabled   bool   `json:"enabled"`
	PublicKey string `json:"public_key"`
	ShortID   string `json:"short_id"`
}

type MultiplexOptions struct {
	Enabled        bool   `json:"enabled"`
	Protocol       string `json:"protocol"`
	MaxConnections int    `json:"max_connections"`
	MinStreams     int    `json:"min_streams"`
	MaxStreams     int    `json:"max_streams"`
	Padding        bool   `json:"padding"`
}

type ListenFields struct {
	Listen                      string `json:"listen"`
	ListenPort                  int    `json:"listen_port"`
	TCPFastOpen                 bool   `json:"tcp_fast_open"`
	UDPFragment                 bool   `json:"udp_fragment"`
	Sniff                       bool   `json:"sniff"`
	SniffOverrideDestination    bool   `json:"sniff_override_destination"`
	SniffTimeout                string `json:"sniff_timeout"`
	DomainStrategy              string `json:"domain_strategy"`
	UDPTimeout                  int    `json:"udp_timeout"`
	ProxyProtocol               bool   `json:"proxy_protocol"`
	ProxyProtocolAcceptNoHeader bool   `json:"proxy_protocol_accept_no_header"`
	Detour                      string `json:"detour"`
}

type _Outbound struct {
	Type string `json:"type"`
	Tag  string `json:"tag,omitempty"`

	TrojanOptions TrojanConfig `json:"-"`
}

type Outbound _Outbound

func (h Outbound) MarshalJSON() ([]byte, error) {
	var v any
	switch h.Type {
	case "trojan":
		v = h.TrojanOptions
	default:
		return nil, errors.New("unknown outbound type")
	}
	return json.Marshal(v)
}

// func (o Outbound) MarshalJSON() ([]byte, error) {
// 	type Alias Outbound
// 	return json.Marshal(&struct {
// 		*Alias
// 	}{
// 		Alias: (*Alias)(&o),
// 	})
// }

func (o *Outbound) UnmarshalJSON(data []byte) error {
	type Alias Outbound
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	return nil
}

func parseTrojanURL(urlStr string) (TrojanConfig, error) {
	var config TrojanConfig
	escapedURL, err := url.QueryUnescape(urlStr)
	var args map[string]string
	if strings.Contains(escapedURL, "trojan://") {
		escapedURL = strings.Replace(escapedURL, "trojan://", "", 1)
		args["password"] = strings.Split(escapedURL, "@")[0]
		for _, v := range strings.Split(strings.Split(escapedURL, "@")[1], "&") {
			args[strings.Split(v, "=")[0]] = strings.Split(v, "=")[1]
		}
		fmt.Println(args)
		fmt.Println(escapedURL)
	} else {
		return config, errors.New("invalid trojan url")
	}
	u, err := url.Parse(urlStr)
	if err != nil {
		return config, err
	}

	query := u.Query()
	// for k, v := range query {
	// 	switch k {
	// 	case "allowInsecure":
	// 		if v[0] == "true" {
	// 			config.TLS.Insecure = true
	// 		}
	// 	case "alpn":
	// 		config.TLS.ALPN = []string{v[0]}
	// 	case "headerType":
	config.Type = "trojan"
	config.Tag = "trojan-out"
	config.Server = u.Hostname()
	config.Port, _ = strconv.Atoi(u.Port())

	config.Password = query.Get("telegram-id-privatevpns")
	config.Network = query.Get("type")

	// Parse TLS options, multiplex options, and transport options if available

	return config, nil
}
func main() {
	// outbound := Outbound{
	// 	Type: "vless",
	// 	Tag:  "vless-out",
	// 	ListenFields: ListenFields{
	// 		Listen:                      "::",
	// 		ListenPort:                  5353,
	// 		TCPFastOpen:                 false,
	// 		UDPFragment:                 false,
	// 		Sniff:                       false,
	// 		SniffOverrideDestination:    false,
	// 		SniffTimeout:                "300ms",
	// 		DomainStrategy:              "prefer_ipv6",
	// 		UDPTimeout:                  300,
	// 		ProxyProtocol:               false,
	// 		ProxyProtocolAcceptNoHeader: false,
	// 		Detour:                      "another-in",
	// 	},
	// 	TLS: TLSOptions{
	// 		Enabled: true,
	// 	},
	// 	Multiplex: MultiplexOptions{
	// 		Enabled:        true,
	// 		Protocol:       "smux",
	// 		MaxConnections: 4,
	// 		MinStreams:     4,
	// 		MaxStreams:     0,
	// 		Padding:        false,
	// 	},
	// }

	// jsonData, err := json.Marshal(outbound)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(string(jsonData))
	c, _ := parseTrojanURL("trojan://telegram-id-privatevpns@13.48.26.241:22222?allowInsecure=&alpn=http%2F1.1&headerType=none&mode=&security=tls&serviceName=&sni=trj.rollingnext.co.uk&type=tcp#🦀PF153926")
	//marshal and save in file
	Outbound := Outbound{Type: "trojan", Tag: "trojan-out", TrojanOptions: c}
	jsonData, err := json.Marshal(Outbound)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}
