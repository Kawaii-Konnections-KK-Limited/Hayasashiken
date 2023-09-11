package main

import (

	// _ "net/http/pprof"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/foreignusage"
)

func main() {

	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	foreignusage.InitService(nil, nil)
}
