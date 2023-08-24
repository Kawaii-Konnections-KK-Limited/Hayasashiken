package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	cmdUtils "github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/cmd/utils"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/run"
)

func init() {
	_, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {

		<-sigCh

		fmt.Println("Received SIGTERM signal")

		cancel() // Cancel the context when SIGINT is received

	}()
}

type Pair struct {
	Ping int32
	Link string
}

// func main() {
// 	fmt.Println(os.Args)
// 	if len(os.Args) < 1 {
// 		fmt.Println("config mode initiated by default, if you want to run by link use: ./v2ray_test link destinationlink")
// 		run.RunByConfigFile()

// 	} else {
// 		// http://cp.cloudflare.com:80 https://icanhazip.com/
// 		res, err := run.SingByLink(&os.Args[1], "https://icanhazip.com/", 50553)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(res)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	res, _ := models.GetAllRecords()
// 	var pairs []Pair

// 	// var pings []int32
// 	for i, v := range res {
// 		link := v.Link
// 		wg.Add(1)
// 		port := i + 50000
// 		go func(link *string, port int) {
// 			defer wg.Done()

// 			r, _ := run.SingByLink(link, "http://cp.cloudflare.com/", port)
// 			pairs = append(pairs, Pair{
// 				Ping: r,
// 				Link: *link,
// 			})
// 		}(&link, port)
// 	}
// 	wg.Wait()
// 	for _, k := range pairs {
// 		fmt.Printf("link: %s RTT: %d \n", k.Link, k.Ping)
// 	}

// }

func main() {

	pairs := cmdUtils.ReadLinksFromFile("path to file")
	var ports []int
	var pp []int32
	var testurl = "https://icanhazip.com/"
	var timeout int32 = 5000
	var baseBroadcast = ""

	for i, v := range pairs {
		link := v

		port := i + 50000

		go func(link *string, port int) {

			r, _ := run.SingByLinkProxy(link, &testurl, &port, &timeout, &baseBroadcast)
			fmt.Println(r)
			pp = append(pp, r)
			if r < 1000 {
				if r != 0 {
					ports = append(ports, port)
				}
			}

		}(&link, port)
	}
	returned := false
	for {

		if len(ports) > 0 && !returned && len(pairs) == len(pp) {

			fmt.Println(ports)
			returned = true
		}

	}

}
