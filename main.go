package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/uussoop/v2ray_test/app"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go app.Run(&wg)
	wg.Wait()
	// time.Sleep(5 * time.Second)
	proxyUrl, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error:", err)
	}
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	if len(os.Args) < 2 {
		fmt.Println("Please provide a link.")
		return
	}
	link := os.Args[1]
	timeout := int32(50000) // timeout in milliseconds
	rtt, testerr := app.UrlTest(client, link, timeout)
	if testerr != nil {
		fmt.Println("Error:", testerr)
	} else {
		fmt.Println("RTT:", rtt, "ms")
	}

}
