package run

import (
	"fmt"
	"sync"

	"github.com/uussoop/v2ray_test/core"
	"github.com/uussoop/v2ray_test/raytest"
)

func RunByConfigFile() (int32, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	go core.Run(&wg)
	wg.Wait()
	// time.Sleep(5 * time.Second)
	res, err := raytest.GetTest("8081", "https://www.google.com")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return res, nil
}
