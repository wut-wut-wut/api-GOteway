package main

import (
	"fmt"
	"net/http"
	"regexp"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.ListenAndServe(":8000", nil)
}

func getMatchingRoute(path string, routes []RouteConfig) *RouteConfig {
	var matchingRoute *RouteConfig = nil
	for i, route := range routes {
		sampleRegex := regexp.MustCompile("^" + route.Path)
		match := sampleRegex.MatchString(path)
		if match {
			fmt.Printf("%s and %s : %t \n", path, route.Path, match)
			matchingRoute = &routes[i]
		}
	}

	return matchingRoute
}

func InitFilters(rc []RouteConfig) []RouteFilter {
	rfs := make([]RouteFilter, 1)

	var rf RouteFilter = StripPrefixFilter{}
	rfs[0] = rf

	return rfs
}
