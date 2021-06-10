package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Routes []RouteConfig
	Server ServerConfig
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

type ServerConfig struct {
	Port     string
	CertFile string `yaml:"cert-file"`
	KeyFile  string `yaml:"key-file"`
}

func GetConfig() Config {

	var configPath string
	flag.StringVar(&configPath, "config", "", "Usage")
	flag.Parse()

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
