package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	cmdUtils "github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/cmd/utils"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/run"
)

// func init() {
// 	_, cancel := context.WithCancel(context.Background())
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh, syscall.SIGINT)
// 	go func() {

// 		<-sigCh

// 		fmt.Println("Received SIGTERM signal")

// 		cancel() // Cancel the context when SIGINT is received

// 	}()
// }

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
var testurl = "https://icanhazip.com/"
var timeout int32 = 10000
var baseBroadcast = "127.0.0.1"
var upperBoundPingLimit int32 = 5000
var ports []int

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {

		<-sigCh

		fmt.Println("Received sigint signal")

		cancel() // Cancel the context when SIGINT is received

	}()

	pairs := cmdUtils.ReadLinksFromFile("cmd/links.txt")
	// var pp []int32

	var counts int = 0
	for i, v := range pairs {
		link := v
		port := i + 50000

		go start(&link, port, ctx, &counts)
	}
	returned := false
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context is done")
			return
		default:

			if !returned && len(pairs) == counts {

				fmt.Println(ports)
				returned = true
			}

			if returned && len(ports) == 0 {
				fmt.Println("all tested nothing works")
				return
			}

		}

	}

}
func start(link *string, port int, ctx context.Context, counts *int) {
	kills := make(chan bool, 1)
	defer close(kills)
	r, _ := run.SingByLinkProxy(link, &testurl, &port, &timeout, &baseBroadcast, ctx, &kills)

	if r < upperBoundPingLimit && r != 0 {

		ports = append(ports, port)

	} else {
		kills <- true

	}

	*counts++

}
