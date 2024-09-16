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
		return cf, err
	}

	defer confFile.Close()

	jsonParser := json.NewDecoder(confFile)
	err = jsonParser.Decode(&cf)

	if err != nil {
		return cf, err
	}

	return cf, nil
}
