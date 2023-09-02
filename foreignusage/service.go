package foreignusage

import (
	"errors"
	"fmt"
	"os"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/foreignusage/api"
)

func InitService() {

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println(errors.New("please set PORT environment variable"))
		port = "8080"
		fmt.Println("Defaulting to port ", port)

	}

	err := api.InitRouter().Run(":" + port)

	if err != nil {
		return
	}

}
