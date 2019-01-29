package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is type for config values
type Config struct {
	Env  string `json:"Env"`
	Port string `json:"Port"`
	DB   struct {
		Name string `json:"Name"`
		URL  string `json:"URL"`
		User string `json:"User"`
		Pass string `json:"Pass"`
	} `json:"DB"`
}

// Values returns config values
func Values() Config {
	jsonFile, err := os.Open("config/config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)

	return config
}
