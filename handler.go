package main

import (
	"log"
	"net/http"
	"regexp"
)

func commonHandler(f []RouteFilter, c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := getMatchingRoute(r.URL.Path, c.Routes)

		if route == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			fc := FilterContext{}
			fc.Filter = &route.Filters[0]
			fc.RequestURI = r.URL.RequestURI()
			fc.TargetURL = route.Url

			for _, filter := range f {
				filter.Process(&fc)
			}

			DoRequest(&fc, w, r)
		}
	}
}

func getMatchingRoute(path string, routes []RouteConfig) *RouteConfig {
	var matchingRoute *RouteConfig = nil
	for i, route := range routes {
		sampleRegex := regexp.MustCompile("^" + route.Path)
		match := sampleRegex.MatchString(path)
		if match {
			log.Printf("%s and %s : %t \n", path, route.Path, match)
			matchingRoute = &routes[i]
		}
	}

	return matchingRoute
}
