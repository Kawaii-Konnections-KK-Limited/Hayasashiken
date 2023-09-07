package foreignusage

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/foreignusage/api"
)

func InitService(certFile, keyFile *string) {

	port := os.Getenv("PORT")
	addr := os.Getenv("ADDRESS")

	if port == "" {
		fmt.Println(errors.New("please set PORT environment variable"))
		port = "8080"
		fmt.Println("Defaulting to port ", port)

	}
	if addr == "" {
		fmt.Println(errors.New("please set ADDRESS environment variable"))
		addr = "0.0.0.0"
		fmt.Println("Defaulting to addr ", addr)
	}
	if addr != "" && certFile != nil && keyFile != nil {
		log.Println("running with tls")
		api.InitRouter().RunTLS(fmt.Sprintf("%s:%s", addr, port), *certFile, *keyFile)
	} else {
		log.Println("running without tls")

		api.InitRouter().Run(fmt.Sprintf("%s:%s", addr, port))
	}

}
