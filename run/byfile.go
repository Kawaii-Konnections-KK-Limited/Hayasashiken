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

	go core.Run(&wg)
	wg.Wait()
	// time.Sleep(5 * time.Second)
	res, err := raytest.GetTest("8081", "https://www.google.com", 5000)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return res, nil
}
