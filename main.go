package main

import (
	"fmt"
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/setup"
	"github.com/willchat-ofc/api-willchat-golang/internal/setup/config"
)

func main() {
	port := ":7070"
	config.LoadEnvFile(".env")

	fmt.Println("server is running with port", port)

	err := http.ListenAndServe(port, setup.Server())

	if err != nil {
		fmt.Println("error ocurred: ", err.Error())
	}
}
