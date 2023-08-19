package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/run"
)

func init() {
	_, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {

		<-sigCh

		fmt.Println("Received SIGINT signal")

		cancel() // Cancel the context when SIGINT is received

	}()
}

type Pair struct {
	Ping int32
	Link string
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 1 {
		fmt.Println("config mode initiated by default, if you want to run by link use: ./v2ray_test link destinationlink")
		run.RunByConfigFile()

	} else {
		// http://cp.cloudflare.com:80 https://icanhazip.com/
		res, err := run.SingByLink(&os.Args[1], "https://icanhazip.com/", 50553)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

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
