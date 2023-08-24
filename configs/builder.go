package configs

import (
	"encoding/json"
	"fmt"
)

func Configbuilder(Rawurl *string, InPort *int, InIp *string) ([]byte, error) {

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
	var logs = log{
		Disabled: true,
	}
	if *InPort == 0 {
		*InPort = 8081
	} else {
		inb.ListenPort = *InPort
	}
	if *InIp != "" {
		inb.Listen = *InIp
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
		Log:       logs,
	}
	configJson, err := config.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return json.RawMessage(configJson), nil
}
