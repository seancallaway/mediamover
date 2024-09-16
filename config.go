/*
layout of config file
{
	"api_key": "",
	"tv_root": "~/Videos/TV",
	"movie_root": "~/Videos/Movies"
}
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ApiKey    string `json:"api_key"`
	TvRoot    string `json:"tv_root"`
	MovieRoot string `json:"movie_root"`
}

func NewConfig(filename string) (Config, error) {
	var cf Config

	confFile, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the config file
			configTemplate := []byte("{\n\t\"api_key\": \"\",\n\t\"tv_root\": \"\",\n\t\"movie_root\": \"\"\n}\n")
			err = os.WriteFile(filename, configTemplate, 0640)
			if err != nil {
				return cf, err
			} else {
				return cf, fmt.Errorf("Config file %s did not exist, was created, and requires configuration", filename)
			}
		} else {
			return cf, err
		}
	}

	defer confFile.Close()

	jsonParser := json.NewDecoder(confFile)
	err = jsonParser.Decode(&cf)

	if err != nil {
		return cf, err
	}

	return cf, nil
}
