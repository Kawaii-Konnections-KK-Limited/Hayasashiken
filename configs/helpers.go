package configs

import (
	"encoding/json"
	"errors"
)

func (m StringStringMap) Get(key string) string {
	return m[key]
}
func networkType(network *string, q *StringStringMap) (any, string) {

	switch *network {
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
	var t []any

	for _, item := range h {
		switch v := item.(type) {
		case MixedInbound:
			t = append(t, v)

		default:
			return nil, errors.New("unknown inbound type")
		}
	}
	return json.Marshal(t)
}

func (h Config) MarshalJSON() ([]byte, error) {

	var t map[string]any
	t = make(map[string]any)
	inb, _ := h.Inbounds.MarshalJSON()
	otb, _ := h.Outbounds.MarshalJSON()
	t["inbounds"] = json.RawMessage(inb)
	t["outbounds"] = json.RawMessage(otb)
	t["route"] = h.Route

	return json.Marshal(t)

}
