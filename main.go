package main

import (
	"log"
	"net/http"
)

type FilterContext struct {
	Filter     *FilterConfig
	R          *http.Request
	TargetURL  string
	RequestURI string
}

type RouteFilter interface {
	Process(f *FilterContext, w http.ResponseWriter, r *http.Request) error
}

func main() {
	c := GetConfig()
	f := InitFilters(c.Routes)

	http.HandleFunc("/", commonHandler(f, c))

	var err error = nil
	if len(c.Server.CertFile) > 0 {
		log.Printf("Starting server on https://localhost:%s", c.Server.Port)
		err = http.ListenAndServeTLS(":"+c.Server.Port, c.Server.CertFile, c.Server.KeyFile, nil)
	} else {
		log.Printf("Starting server on http://localhost:%s", c.Server.Port)
		err = http.ListenAndServe(":"+c.Server.Port, nil)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func InitFilters(rc []RouteConfig) map[string]RouteFilter {
	rfs := make(map[string]RouteFilter)

	rfs["StripPrefixFilter"] = StripPrefixFilter{}
	rfs["BasicAuthFilter"] = BasicAuthFilter{}

	return rfs
}
