package run

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/uussoop/v2ray_test/configs"
	"github.com/uussoop/v2ray_test/core"
	"github.com/uussoop/v2ray_test/raytest"
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
