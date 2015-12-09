package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type configuration struct {
	Port      int
	Redirects map[string]redirect
}

type redirect struct {
	Host string
	Code int
}

func getConfig(fileName string) (*configuration, bool) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	var config configuration
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	return &config, true
}

func (config *configuration) getRedirect(host string) (*redirect, bool) {
	if redirect, ok := config.Redirects[host]; ok {
		return &redirect, ok
	}
	if redirect, ok := config.Redirects["*"]; ok {
		return &redirect, ok
	}
	return nil, false
}
