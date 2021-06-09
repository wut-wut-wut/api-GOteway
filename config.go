package main

import (
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