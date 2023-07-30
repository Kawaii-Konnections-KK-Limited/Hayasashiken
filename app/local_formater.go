package app

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/sagernet/sing-box/option"
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
	Type           string           `json:"type"`
	Tag            string           `json:"tag"`
	Server         string           `json:"server"`
	ServerPort     int              `json:"server_port"`
	Security       string           `json:"security"`
	AlterID        int              `json:"alter_id"`
	GlobalPadding  bool             `json:"global_padding"`
	Id             string           `json:"uuid"`
	AuthLen        bool             `json:"authenticated_length"`
	Network        string           `json:"network"`
	TLS            TLSOptions       `json:"tls"`
	PacketEncoding string           `json:"packet_encoding"`
	Multiplex      MultiplexOptions `json:"multiplex"`
	Transport      TransportOptions `json:"transport"`
}

type VLESSConfig struct {
	Type           string           `json:"type"`
	Tag            string           `json:"tag"`
	Server         string           `json:"server"`
	ServerPort     int              `json:"server_port"`
	Id             string           `json:"uuid"`
	Network        string           `json:"network"`
	TLS            TLSOptions       `json:"tls"`
	PacketEncoding string           `json:"packet_encoding"`
	Flow           string           `json:"flow"`
	Multiplex      MultiplexOptions `json:"multiplex"`
	Transport      TransportOptions `json:"transport"`
}

// HTTP
type HTTPConfig struct {
	Type        string            `json:"type"`
	Host        []string          `json:"host"`
	Path        string            `json:"path"`
	Method      string            `json:"method"`
	Headers     map[string]string `json:"headers"`
	IdleTimeout string            `json:"idle_timeout"`
	PingTimeout string            `json:"ping_timeout"`
}

// WebSocket
type WSConfig struct {
	Type            string            `json:"type"`
	Path            string            `json:"path"`
	Headers         map[string]string `json:"headers"`
	MaxEarlyData    int               `json:"max_early_data"`
	EarlyDataHeader string            `json:"early_data_header_name"`
}

// QUIC
type QUICConfig struct {
	Type string `json:"type"`
}

// gRPC
type GRPCConfig struct {
	Type                string `json:"type"`
	ServiceName         string `json:"service_name"`
	IdleTimeout         string `json:"idle_timeout"`
	PingTimeout         string `json:"ping_timeout"`
	PermitWithoutStream bool   `json:"permit_without_stream"`
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

type MixedInbound struct {
	Type                        string `json:"type"`
	Tag                         string `json:"tag"`
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
	SetSystemProxy              bool   `json:"set_system_proxy"`
}
type Route interface{}
type _Outbound []any
type _Inbound []any
type Config struct {
	Inbounds  _Inbound  `json:"inbounds"`
	Outbounds _Outbound `json:"outbounds"`
	Route     Route     `json:"route"`
}
type ConfigTypes struct {
	Vless  VLESSConfig  `json:"vless"`
	Vmess  VMessConfig  `json:"vmess"`
	Trojan TrojanConfig `json:"trojan"`
}
type StringStringMap map[string]string

func (m StringStringMap) Get(key string) string {
	return m[key]
}
func networkType(network string, q *StringStringMap) (any, string) {

	switch network {
	case "ws":
		return WSConfig{
			Type: "ws",
			Path: q.Get("path"),
		}, "tcp"
	case "grpc":
		return GRPCConfig{
			Type:        "grpc",
			ServiceName: q.Get("serviceName"),
			IdleTimeout: "3s",
			PingTimeout: "10s",
		}, "tcp"
	case "http":
		return HTTPConfig{
			Type:   "http",
			Path:   q.Get("path"),
			Method: q.Get("method"),
		}, "tcp"
	case "quic":
		return QUICConfig{
			Type: "quic",
		}, "udp"

	case "tcp":
		return nil, "tcp"
	case "udp":
		return nil, "udp"
	default:
		return nil, "tcp"
	}

}

func (h _Outbound) MarshalJSON() ([]byte, error) {
	var t []any
	for _, item := range h {
		switch v := item.(type) {
		case TrojanConfig:
			t = append(t, v)
		case VMessConfig:
			t = append(t, v)

		case VLESSConfig:
			t = append(t, v)

		default:
			return nil, errors.New("unknown outbound type")
		}
	}
	return json.Marshal(t)
}
func (h _Inbound) MarshalJSON() ([]byte, error) {
	for _, item := range h {
		switch v := item.(type) {
		case MixedInbound:
			return json.Marshal(v)
		default:
			return nil, errors.New("unknown inbound type")
		}
	}
	return []byte(""), nil
}

func (h Config) MarshalJSON() ([]byte, error) {

	var t map[string]any
	t = make(map[string]any)
	inb, _ := h.Inbounds.MarshalJSON()
	otb, _ := h.Outbounds.MarshalJSON()
	t["inbounds"] = json.RawMessage(inb)
	t["outbounds"] = json.RawMessage(otb)
	t["routing"] = h.Route

	return json.Marshal(t)

}
func ParseTrojanURL(rawURL string) (TrojanConfig, error) {

	// Set default values
	cfg := TrojanConfig{
		Type: "trojan",
		Tag:  "trojan-out",
		TLS: TLSOptions{
			Enabled: true,
		},
	}

	u, _ := url.Parse(rawURL)

	// Parse userinfo for password
	userInfo := strings.Split(u.User.Username(), "@")
	if len(userInfo) > 0 {
		cfg.Password = userInfo[0]
	}

	// Parse hostname
	if u.Hostname() != "" {
		cfg.Server = u.Hostname()
	}

	// Parse port
	if u.Port() != "" {
		cfg.Port, _ = strconv.Atoi(u.Port())
	}

	// Parse query parameters
	if alpn := u.Query().Get("alpn"); alpn != "" {
		cfg.TLS.ALPN = strings.Split(alpn, ",")
	}

	if sni := u.Query().Get("sni"); sni != "" {
		cfg.TLS.ServerName = sni
	}
	if networkType := u.Query().Get("type"); networkType != "" {
		cfg.Network = networkType
	}

	return cfg, nil
}

func ParseVmessURL(rawURL string) (VMessConfig, error) {

	// Decode base64 string
	data, err := base64.RawURLEncoding.DecodeString(strings.Split(rawURL, "//")[1])

	if err != nil {
		fmt.Println(err)
	}
	//convert data to map

	mapped := make(StringStringMap)
	json.Unmarshal(data, &mapped)
	fmt.Println(mapped)
	var cfg VMessConfig
	cfg.Type = "vmess"
	cfg.Tag = "vmess-out"
	if server := mapped["add"]; server != "" {
		cfg.Server = server
	}
	if port := mapped["port"]; port != "" {
		cfg.ServerPort, _ = strconv.Atoi(port)
	}
	if security := mapped["scy"]; security != "" {
		cfg.Security = security
	}
	if alterID := mapped["aid"]; alterID != "" {
		cfg.AlterID, _ = strconv.Atoi(alterID)
	}
	if network := mapped["net"]; network != "" {

		cfg.Transport, cfg.Network = networkType(network, &mapped)

	}
	if tls := mapped["tls"]; tls != "" {
		cfg.TLS.Enabled = true
	}
	if sni := mapped["sni"]; sni != "" {
		cfg.TLS.ServerName = sni
	}
	if id := mapped["id"]; id != "" {
		cfg.Id = id
	}

	// Unmarshal JSON into struct

	return cfg, nil
}

func ParseVlessUrl(rawURL string) (VLESSConfig, error) {

	u, _ := url.Parse(rawURL)
	port, _ := strconv.Atoi(u.Port())
	config := VLESSConfig{
		Type:       "vless",
		Tag:        "vless-out",
		Server:     u.Hostname(),
		ServerPort: port,
		Id:         strings.TrimPrefix(u.User.Username(), "@"),
	}

	q := u.Query()
	m := make(StringStringMap)
	for k, v := range q {
		m[k] = v[0]
	}

	if sni := q.Get("sni"); sni != "" {
		config.TLS.ServerName = sni
	}

	if network := q.Get("type"); network != "" {
		config.Transport, config.Network = networkType(network, &m)

	}
	if flow := q.Get("flow"); flow != "" {
		config.Flow = flow
	}
	if security := q.Get("security"); security != "" {
		if security == "reality" {
			config.TLS.Reality.Enabled = true
			config.TLS.Reality.PublicKey = q.Get("publicKey")
			config.TLS.Reality.ShortID = q.Get("shortId")
		}
	}
	if fingerprint := q.Get("fp"); fingerprint != "" {
		config.TLS.UTLS.Enabled = true
		config.TLS.UTLS.Fingerprint = fingerprint
	}

	// Parse other params like security etc

	return config, nil
}

func main() {

	var outbound _Outbound
	c, _ := ParseTrojanURL("trojan://telegram-id-privatevpns@13.48.26.241:22222?allowInsecure=&alpn=http%2F1.1&headerType=none&mode=&security=tls&serviceName=&sni=trj.rollingnext.co.uk&type=tcp#🦀PF153926")
	outbound = append(outbound, c)

	// c2, _ := ParseVmessURL("vmess://eyJhaWQiOiAiMCIsICJhbHBuIjogIiIsICJob3N0IjogInBybmV3c3dpcmUuY29tIiwgIm5ldCI6ICJ3cyIsICJwYXRoIjogIi8iLCAicG9ydCI6ICIzNTY0MyIsICJzY3kiOiAiIiwgInNuaSI6ICIiLCAidGxzIjogIm5vbmUiLCAidHlwZSI6ICJub25lIiwgInYiOiAiMiIsICJwcyI6ICJcdWQ4M2VcdWRkODRQRjE1Mzk2NyIsICJhZGQiOiAiMTg1LjIwNi45My40IiwgImlkIjogImE4NDllNmRhLWYzMjQtNGI1NS05ODI3LWQ2OGRhZWVjMmE5MCJ9")

	// outbound = append(outbound, c2)
	// c3, _ := ParseVlessUrl("vless://d39f7166-9824-4064-b69d-e15a26ddd5a8@join.mdvpnsec.cfd:2096?encryption=&flow=&headerType=&host=&path=&security=reality&sni=www.speedtest.net&type=grpc#🚀PF167273")

	// outbound = append(outbound, c3)
	// test, _ := outbound.MarshalJSON()
	// println(string(test))

	// outboundJson, _ := outbound.MarshalJSON()
	// println(string(outboundJson))

	port := 8081
	var inbound _Inbound
	inbound = append(inbound, MixedInbound{
		Type:                        "mixed",
		Tag:                         "mixed-in",
		Listen:                      "127.0.0.1",
		ListenPort:                  port,
		TCPFastOpen:                 false,
		UDPFragment:                 false,
		Sniff:                       false,
		SniffOverrideDestination:    false,
		SniffTimeout:                "300ms",
		UDPTimeout:                  300,
		ProxyProtocol:               false,
		ProxyProtocolAcceptNoHeader: false,
		SetSystemProxy:              false,
	})
	jin, _ := inbound.MarshalJSON()
	fmt.Println(string(jin))
	config := Config{
		Inbounds:  inbound,
		Outbounds: outbound,
		Route:     nil,
	}
	configJson, _ := config.MarshalJSON()
	println(string(configJson))
	// var (
	// 	configContent []byte
	// 	err           error
	// )
	// if err != nil {
	// 	fmt.Printf("error")
	// }
	// configContent, err = os.ReadFile("config.json")
	// if err != nil {
	// 	fmt.Printf("file doesnt exist")
	// }
	var options option.Options
	err := options.UnmarshalJSON(configJson)
	if err != nil {
		fmt.Println(err)
	}

}
