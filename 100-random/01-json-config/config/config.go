package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Env      string `json:"Env"`
	Port     int    `json:"Port"`
	Database struct {
		DBname string `json:"DBname"`
		DBurl  string `json:"DBurl"`
		DBuser string `json:"DBuser"`
		DBpass string `json:"DBpass"`
	} `json:"Database"`
}

func init() {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)

	fmt.Println(config)
}
