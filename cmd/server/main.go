package main

import (
	"fmt"

	"github.com/brunacotrim/api-product/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
