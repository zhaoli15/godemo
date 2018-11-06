package main

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	DbSource string `json:"db_source"`
}

func initConfigFromFile(file string) (*Config, error) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var flag Config
	if err := json.Unmarshal(data, &flag); err != nil {
		return nil, err
	}
	return &flag, nil
}
