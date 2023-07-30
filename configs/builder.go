package configs

import (
	"encoding/json"
	"fmt"
)

var inb = MixedInbound{
	Type:                        "mixed",
	Tag:                         "mixed-in",
	Listen:                      "127.0.0.1",
	ListenPort:                  8081,
	TCPFastOpen:                 false,
	UDPFragment:                 false,
	Sniff:                       false,
	SniffOverrideDestination:    false,
	SniffTimeout:                "300ms",
	UDPTimeout:                  300,
	ProxyProtocol:               false,
	ProxyProtocolAcceptNoHeader: false,
	SetSystemProxy:              false,
}

func Configbuilder(Rawurl *string, InPort int) ([]byte, error) {
	if InPort == 0 {
		InPort = 8081
		inb.ListenPort = InPort
	}
	var outbound _Outbound
	c, err := parseUrl(Rawurl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	outbound = append(outbound, c)
	var inbound _Inbound
	inbound = append(inbound, inb)
	config := Config{
		Inbounds:  inbound,
		Outbounds: outbound,
		Route:     nil,
	}
	configJson, err := config.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(configJson))
	return json.RawMessage(configJson), nil
}
