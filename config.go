package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Routes []RouteConfig
}

type RouteConfig struct {
	Name    string
	Path    string
	Url     string
	Filters []FilterConfig
}

type FilterConfig struct {
	Name       string
	Properties map[string]string
}

func GetConfig() Config {

	var configPath string
	flag.StringVar(&configPath, "config", "", "Usage")
	flag.Parse()
	fmt.Println(configPath)

	c := Config{}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("File configuration %s loaded successfully", configPath)

	return c
}
