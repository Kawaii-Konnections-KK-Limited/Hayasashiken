package run

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/configs"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/core"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/raytest"
)

func SingByLink(Rawurl *string, Testurl string, InputPort int) (int32, error) {

	c, err := configs.Configbuilder(Rawurl, InputPort)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var wg sync.WaitGroup
	wg.Add(1)

	go core.RunByLink(&wg, &c)
	wg.Wait()
	res, err := raytest.GetTest(strconv.Itoa(InputPort), Testurl)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return res, nil

}
