package configs

type LogObject struct {
	AccessLog string `json:"access"`
	ErrorLog  string `json:"error"`
	LogLevel  string `json:"loglevel"`
	DnsLog    bool   `json:"dnsLog"`
}

type ApiObject struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}

type DnsObject struct {
	Hosts                  map[string]interface{} `json:"hosts"`
	Servers                []interface{}          `json:"servers"`
	ClientIP               string                 `json:"clientIp"`
	QueryStrategy          string                 `json:"queryStrategy"`
	DisableCache           bool                   `json:"disableCache"`
	DisableFallback        bool                   `json:"disableFallback"`
	DisableFallbackIfMatch bool                   `json:"disableFallbackIfMatch"`
	Tag                    string                 `json:"tag"`
}

type RoutingObject struct {
	DomainStrategy string           `json:"domainStrategy"`
	DomainMatcher  string           `json:"domainMatcher"`
	Rules          []RuleObject     `json:"rules"`
	Balancers      []BalancerObject `json:"balancers"`
}

type RuleObject struct {
	DomainMatcher string            `json:"domainMatcher"`
	Type          string            `json:"type"`
	Domain        []string          `json:"domain"`
	IP            []string          `json:"ip"`
	Port          string            `json:"port"`
	SourcePort    string            `json:"sourcePort"`
	Network       string            `json:"network"`
	Source        []string          `json:"source"`
	User          []string          `json:"user"`
	InboundTag    []string          `json:"inboundTag"`
	Protocol      []string          `json:"protocol"`
	Attrs         map[string]string `json:"attrs"`
	OutboundTag   string            `json:"outboundTag"`
	BalancerTag   string            `json:"balancerTag"`
}

type BalancerObject struct {
	Tag      string   `json:"tag"`
	Selector []string `json:"selector"`
}

type PolicyObject struct {
	Levels map[string]LevelPolicyObject `json:"levels"`
	System SystemPolicyObject           `json:"system"`
}

type LevelPolicyObject struct {
	Handshake         int  `json:"handshake"`
	ConnIdle          int  `json:"connIdle"`
	UplinkOnly        int  `json:"uplinkOnly"`
	DownlinkOnly      int  `json:"downlinkOnly"`
	StatsUserUplink   bool `json:"statsUserUplink"`
	StatsUserDownlink bool `json:"statsUserDownlink"`
	BufferSize        int  `json:"bufferSize"`
}

type SystemPolicyObject struct {
	StatsInboundUplink    bool `json:"statsInboundUplink"`
	StatsInboundDownlink  bool `json:"statsInboundDownlink"`
	StatsOutboundUplink   bool `json:"statsOutboundUplink"`
	StatsOutboundDownlink bool `json:"statsOutboundDownlink"`
}

type StatsObject struct {
}

type ReverseObject struct {
	Bridges []BridgeObject `json:"bridges"`
	Portals []PortalObject `json:"portals"`
}

type BridgeObject struct {
	Tag    string `json:"tag"`
	Domain string `json:"domain"`
}

type PortalObject struct {
	Tag    string `json:"tag"`
	Domain string `json:"domain"`
}

type FakeDnsObject struct {
	IpPool   string `json:"ipPool"`
	PoolSize int    `json:"poolSize"`
}

type InboundObject struct {
	Listen         string               `json:"listen"`
	Port           interface{}          `json:"port"`
	Protocol       string               `json:"protocol"`
	Settings       interface{}          `json:"settings"`
	StreamSettings StreamSettingsObject `json:"streamSettings"`
	Tag            string               `json:"tag"`
	Sniffing       SniffingObject       `json:"sniffing"`
	Allocate       AllocateObject       `json:"allocate"`
}

type SniffingObject struct {
	Enabled         bool     `json:"enabled"`
	DestOverride    []string `json:"destOverride"`
	MetadataOnly    bool     `json:"metadataOnly"`
	DomainsExcluded []string `json:"domainsExcluded"`
	RouteOnly       bool     `json:"routeOnly"`
}

type AllocateObject struct {
	Strategy    string `json:"strategy"`
	Refresh     int    `json:"refresh"`
	Concurrency int    `json:"concurrency"`
}

type OutboundObject struct {
	SendThrough    string               `json:"sendThrough"`
	Protocol       string               `json:"protocol"`
	Settings       interface{}          `json:"settings"`
	Tag            string               `json:"tag"`
	StreamSettings StreamSettingsObject `json:"streamSettings"`
	ProxySettings  ProxySettingsObject  `json:"proxySettings"`
	Mux            MuxObject            `json:"mux"`
}

type ProxySettingsObject struct {
	Tag string `json:"tag"`
}

type MuxObject struct {
	Enabled     bool `json:"enabled"`
	Concurrency int  `json:"concurrency"`
}

type TransportObject struct {
	TcpSettings  TcpObject       `json:"tcpSettings"`
	KcpSettings  KcpObject       `json:"kcpSettings"`
	WsSettings   WebSocketObject `json:"wsSettings"`
	HttpSettings HttpObject      `json:"httpSettings"`
	QuicSettings QuicObject      `json:"quicSettings"`
	GrpcSettings GrpcObject      `json:"grpcSettings"`
}

type TcpObject struct {
	AcceptProxyProtocol bool         `json:"acceptProxyProtocol"`
	Header              HeaderObject `json:"header"`
}

type HeaderObject struct {
	Type string `json:"type"`
}

type KcpObject struct {
	Mtu              int          `json:"mtu"`
	Tti              int          `json:"tti"`
	UplinkCapacity   int          `json:"uplinkCapacity"`
	DownlinkCapacity int          `json:"downlinkCapacity"`
	Congestion       bool         `json:"congestion"`
	ReadBufferSize   int          `json:"readBufferSize"`
	WriteBufferSize  int          `json:"writeBufferSize"`
	Header           HeaderObject `json:"header"`
	Seed             string       `json:"seed"`
}

type WebSocketObject struct {
	AcceptProxyProtocol bool              `json:"acceptProxyProtocol"`
	Path                string            `json:"path"`
	Headers             map[string]string `json:"headers"`
}

type HttpObject struct {
	Host               []string            `json:"host"`
	Path               string              `json:"path"`
	ReadIdleTimeout    int                 `json:"read_idle_timeout"`
	HealthCheckTimeout int                 `json:"health_check_timeout"`
	Method             string              `json:"method"`
	Headers            map[string][]string `json:"headers"`
}

type QuicObject struct {
	Security string       `json:"security"`
	Key      string       `json:"key"`
	Header   HeaderObject `json:"header"`
}

type GrpcObject struct {
	ServiceName         string `json:"serviceName"`
	MultiMode           bool   `json:"multiMode"`
	IdleTimeout         int    `json:"idle_timeout"`
	HealthCheckTimeout  int    `json:"health_check_timeout"`
	PermitWithoutStream bool   `json:"permitWithoutStream"`
	InitialWindowsSize  int    `json:"initial_windows_size"`
}

type StreamSettingsObject struct {
	Network      string          `json:"network"`
	Security     string          `json:"security"`
	XtlsSettings XtlsSettings    `json:"xtlsSettings"`
	TlsSettings  TlsSettings     `json:"tlsSettings"`
	TcpSettings  TcpObject       `json:"tcpSettings"`
	KcpSettings  KcpObject       `json:"kcpSettings"`
	WsSettings   WebSocketObject `json:"wsSettings"`
	HttpSettings HttpObject      `json:"httpSettings"`
	QuicSettings QuicObject      `json:"quicSettings"`
	GrpcSettings GrpcObject      `json:"grpcSettings"`
	Sockopt      SockoptObject   `json:"sockopt"`
}

type XtlsSettings struct {
	ServerName   string              `json:"serverName"`
	Alpn         []string            `json:"alpn"`
	Certificates []CertificateObject `json:"certificates"`
}

type TlsSettings struct {
	ServerName   string              `json:"serverName"`
	Alpn         []string            `json:"alpn"`
	Certificates []CertificateObject `json:"certificates"`
}

type CertificateObject struct {
	OcspStapling    int      `json:"ocspStapling"`
	CertificateFile string   `json:"certificateFile"`
	Certificate     []string `json:"certificate"`
	KeyFile         string   `json:"keyFile"`
	Key             []string `json:"key"`
}

type SockoptObject struct {
	Mark                 int         `json:"mark"`
	TcpFastOpen          interface{} `json:"tcpFastOpen"`
	Tproxy               string      `json:"tproxy"`
	DomainStrategy       string      `json:"domainStrategy"`
	DialerProxy          string      `json:"dialerProxy"`
	AcceptProxyProtocol  bool        `json:"acceptProxyProtocol"`
	TcpKeepAliveInterval int         `json:"tcpKeepAliveInterval"`
}
