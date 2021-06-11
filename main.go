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
	Process(f *FilterContext) bool
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

func InitFilters(rc []RouteConfig) []RouteFilter {
	rfs := make([]RouteFilter, 1)

	var rf RouteFilter = StripPrefixFilter{}
	rfs[0] = rf

	return rfs
}
