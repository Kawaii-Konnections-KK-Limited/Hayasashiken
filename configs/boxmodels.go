package configs

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
	Type        string              `json:"type"`
	Host        []string            `json:"host"`
	Path        string              `json:"path"`
	Method      string              `json:"method"`
	Headers     map[string][]string `json:"headers"`
	IdleTimeout string              `json:"idle_timeout"`
	PingTimeout string              `json:"ping_timeout"`
}
type TCPConfig struct {
	Type string `json:"type"`

	Path    string              `json:"path"`
	Method  string              `json:"method"`
	Headers map[string][]string `json:"headers"`
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
type log struct {
	Disabled  bool   `json:"disabled"`
	Level     string `json:"level"`
	Output    string `json:"output"`
	Timestamp bool   `json:"timestamp"`
}
type Config struct {
	Inbounds  _Inbound  `json:"inbounds"`
	Outbounds _Outbound `json:"outbounds"`
	Route     Route     `json:"route"`
	Log       log       `json:"log"`
}
type ConfigTypes struct {
	Vless  VLESSConfig  `json:"vless"`
	Vmess  VMessConfig  `json:"vmess"`
	Trojan TrojanConfig `json:"trojan"`
}
type StringStringMap map[string]string
