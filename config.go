package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	config map[string]interface{}
)

func InitConfig() (err error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return
}

func GetPluginConfig(id string, key string) interface{} {
	if plg, ok := config["plugins"]; ok {
		if plgid, ok := plg.(map[string]interface{})[id]; ok {
			if val, ok := plgid.(map[string]interface{})[key]; ok {
				return val
			}
		}
	}
	return nil
}