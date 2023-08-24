package run

import (
	"fmt"
	"sync"
	"time"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/configs"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/core"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/raytest"
)

func SingByLink(Rawurl *string, Testurl *string, InputPort *int, TimeOut *int32, InIp *string) (int32, error) {

	c, err := configs.Configbuilder(Rawurl, InputPort, InIp)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var wg sync.WaitGroup
	wg.Add(1)

	go core.RunByLink(&wg, &c)
	wg.Wait()
	res, err := raytest.GetTest(InputPort, Testurl, TimeOut)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return res, nil

}
func SingByLinkProxy(Rawurl *string, Testurl *string, InputPort *int, TimeOut *int32, InIp *string) (int32, error) {

	c, err := configs.Configbuilder(Rawurl, InputPort, InIp)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	instanceReady := make(chan bool, 1)

	go core.RunByLinkProxy(&instanceReady, &c)
	for {
		if <-instanceReady {
			fmt.Println("Instance is ready")
			break
		} else {

			fmt.Println("Instance is not ready")
			time.Sleep(5 * time.Second)
		}
	}
	// wg.Wait()
	res, err := raytest.GetTest(InputPort, Testurl, TimeOut)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return res, nil

}
