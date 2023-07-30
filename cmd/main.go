package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/uussoop/v2ray_test/models"
	"github.com/uussoop/v2ray_test/run"
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

// func main() {
// 	fmt.Println(os.Args)
// 	if len(os.Args) < 3 {
// 		fmt.Println("config mode initiated by default, if you want to run by link use: ./v2ray_test link destinationlink")
// 		run.RunByConfigFile()

//		} else {
//			run.SingByLink(&os.Args[1], os.Args[2])
//		}
//	}
func main() {
	var wg sync.WaitGroup
	res, _ := models.GetAllRecords()
	fmt.Println(len(res))
	for i, v := range res[4:5] {
		wg.Add(1)
		port := i + 50000
		go func(link *string, port int) {
			defer wg.Done()

			_, err := run.SingByLink(link, "https://claude.ai", port)
			fmt.Println(err)
		}(&v.Link, port)
	}
	wg.Wait()
}
