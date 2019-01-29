package main

import (
	"fmt"
	"go-demos-2018/100-random/01-json-config/config"
)

func main() {
	config := config.Values()
	fmt.Printf("Config: Env=%s Port=%s DB.Name=%s\n", config.Env, config.Port, config.DB.Name)
}
