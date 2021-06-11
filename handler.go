package main

import (
	"log"
	"net/http"
	"regexp"
)

func commonHandler(f map[string]RouteFilter, c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := getMatchingRoute(r.URL.Path, c.Routes)

		if route == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			fc := FilterContext{}
			fc.RequestURI = r.URL.RequestURI()
			fc.TargetURL = route.Url

			var err error = nil

			// Apply all the configured filters associate to the Route
			for _, filterConfig := range route.Filters {
				filter, ok := f[filterConfig.Name]
				if ok {
					fc.Filter = &filterConfig
					err = filter.Process(&fc, w, r)

					// Dont process the other filters if last one returned an error
					if err != nil {
						break
					}
				}
			}

			if err == nil {
				DoRequest(&fc, w, r)
			}
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
