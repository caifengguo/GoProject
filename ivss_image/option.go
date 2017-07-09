package main

import (
	"encoding/xml"
	"io/ioutil"
)

type ServerConfig struct {
	ListenUrl string `xml:"locale.listen.url"`
}

var G_Serverconfig ServerConfig

func G_Option_Init(file string) bool {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return false
	}
	err = xml.Unmarshal(content, &G_Serverconfig)
	if err != nil {
		return false
	}
	return true
}
