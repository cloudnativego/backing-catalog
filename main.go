package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/backing-catalog/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("CF Environment not detected.")
	}

	server := service.NewServerFromCFEnv(appEnv)
	server.Run(":" + port)
}
