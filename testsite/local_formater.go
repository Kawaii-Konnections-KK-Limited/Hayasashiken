package app

// func main() {

// 	var outbound _Outbound
// 	c, _ := ParseTrojanURL("trojan://telegram-id-privatevpns@13.48.26.241:22222?allowInsecure=&alpn=http%2F1.1&headerType=none&mode=&security=tls&serviceName=&sni=trj.rollingnext.co.uk&type=tcp#🦀PF153926")
// 	outbound = append(outbound, c)

// 	// c2, _ := ParseVmessURL("vmess://eyJhaWQiOiAiMCIsICJhbHBuIjogIiIsICJob3N0IjogInBybmV3c3dpcmUuY29tIiwgIm5ldCI6ICJ3cyIsICJwYXRoIjogIi8iLCAicG9ydCI6ICIzNTY0MyIsICJzY3kiOiAiIiwgInNuaSI6ICIiLCAidGxzIjogIm5vbmUiLCAidHlwZSI6ICJub25lIiwgInYiOiAiMiIsICJwcyI6ICJcdWQ4M2VcdWRkODRQRjE1Mzk2NyIsICJhZGQiOiAiMTg1LjIwNi45My40IiwgImlkIjogImE4NDllNmRhLWYzMjQtNGI1NS05ODI3LWQ2OGRhZWVjMmE5MCJ9")

// 	// outbound = append(outbound, c2)
// 	// c3, _ := ParseVlessUrl("vless://d39f7166-9824-4064-b69d-e15a26ddd5a8@join.mdvpnsec.cfd:2096?encryption=&flow=&headerType=&host=&path=&security=reality&sni=www.speedtest.net&type=grpc#🚀PF167273")

// 	// outbound = append(outbound, c3)
// 	// test, _ := outbound.MarshalJSON()
// 	// println(string(test))

// 	// outboundJson, _ := outbound.MarshalJSON()
// 	// println(string(outboundJson))

// 	port := 8081
// 	var inbound _Inbound
// 	inbound = append(inbound, MixedInbound{
// 		Type:                        "mixed",
// 		Tag:                         "mixed-in",
// 		Listen:                      "127.0.0.1",
// 		ListenPort:                  port,
// 		TCPFastOpen:                 false,
// 		UDPFragment:                 false,
// 		Sniff:                       false,
// 		SniffOverrideDestination:    false,
// 		SniffTimeout:                "300ms",
// 		UDPTimeout:                  300,
// 		ProxyProtocol:               false,
// 		ProxyProtocolAcceptNoHeader: false,
// 		SetSystemProxy:              false,
// 	})
// 	jin, _ := inbound.MarshalJSON()
// 	fmt.Println(string(jin))
// 	config := Config{
// 		Inbounds:  inbound,
// 		Outbounds: outbound,
// 		Route:     nil,
// 	}
// 	configJson, _ := config.MarshalJSON()
// 	println(string(configJson))
// 	// var (
// 	// 	configContent []byte
// 	// 	err           error
// 	// )
// 	// if err != nil {
// 	// 	fmt.Printf("error")
// 	// }
// 	// configContent, err = os.ReadFile("config.json")
// 	// if err != nil {
// 	// 	fmt.Printf("file doesnt exist")
// 	// }
// 	var options option.Options
// 	err := options.UnmarshalJSON(configJson)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }
