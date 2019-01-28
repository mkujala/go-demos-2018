package main

import (
	"fmt"
	"go-demos-2018/100-random/01-json-config/config"
)

func main() {
	config := config.Values()
	fmt.Println(config.Port)
}
