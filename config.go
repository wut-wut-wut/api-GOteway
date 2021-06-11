package main

import (
<<<<<<< HEAD
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
=======
    "fmt"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Routes []RouteConfig
}

type RouteConfig struct {
    Name string
    Path string
    Url string
    Filters []FilterConfig
}

type FilterConfig struct {
    Name string
    Properties map[string]string
}

func GetConfig() Config{
    c := Config{}
    var data = `
    routes:
        - name: cars
          path: /cars/*
          url: https://localhost:8080
          filters:
            - name: stripPrefix
              properties:
                depth: 1
        - name: houses
          path: /houses/*
          url: http://localhost:8001
    `
    err := yaml.Unmarshal([]byte(data), &c)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("--- c:\n%v\n\n", c)

    return c
}
>>>>>>> 68f7fd6346f1a0fa3e020d287e27fc2341237346
