package configs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func parseTrojanURL(rawURL *string) (TrojanConfig, error) {

	// Set default values
	cfg := TrojanConfig{
		Type: "trojan",
		Tag:  "trojan-out",
		TLS: TLSOptions{
			Enabled: true,
		},
	}

	u, _ := url.Parse(*rawURL)

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

func parseVmessURL(rawURL *string) (VMessConfig, error) {

	// Decode base64 string
	data, err := base64.RawURLEncoding.DecodeString(strings.Split(*rawURL, "//")[1])

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

		cfg.Transport, cfg.Network = networkType(&network, &mapped)

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

func parseVlessUrl(rawURL *string) (VLESSConfig, error) {

	u, _ := url.Parse(*rawURL)
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
		config.Transport, config.Network = networkType(&network, &m)

	}
	if flow := q.Get("flow"); flow != "" {
		config.Flow = flow
	}
	if security := q.Get("security"); security != "" {
		if security == "reality" {
			config.TLS.Enabled = true
			config.TLS.Reality.Enabled = true
			config.TLS.Reality.PublicKey = q.Get("pbk")
			config.TLS.Reality.ShortID = q.Get("sid")
		}
	}
	if fingerprint := q.Get("fp"); fingerprint != "" {
		config.TLS.UTLS.Enabled = true
		config.TLS.UTLS.Fingerprint = fingerprint
	}

	// Parse other params like security etc

	return config, nil
}

func parseUrl(rawURL *string) (any, error) {
	u, _ := url.Parse(*rawURL)
	switch u.Scheme {
	case "trojan":
		return parseTrojanURL(rawURL)
	case "vmess":
		return parseVmessURL(rawURL)
	case "vless":
		return parseVlessUrl(rawURL)
	default:
		return nil, fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}
}
