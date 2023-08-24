package run

import (
	"fmt"
	"sync"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/core"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/raytest"
)

func RunByConfigFile() (int32, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	var port int = 8081
	testurl := "https://www.google.com"
	var timeout int32 = 5000
	go core.Run(&wg)
	wg.Wait()
	// time.Sleep(5 * time.Second)
	res, err := raytest.GetTest(&port, &testurl, &timeout)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return res, nil
}
