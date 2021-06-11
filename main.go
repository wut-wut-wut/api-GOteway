package main

import (
<<<<<<< HEAD
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
=======
    "fmt"
    "net/http"
    "regexp"
)

type FilterContext struct {
    Filter *FilterConfig
    R *http.Request
    TargetURL string
    RequestURI string
}

type RouteFilter interface {
    Process(f *FilterContext) bool
}

func main() {
    c := GetConfig()
    f := InitFilters(c.Routes)
    
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        route := getMatchingRoute(r.URL.Path, c.Routes)
          
        if route == nil {
            http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
        } else {
             //fmt.Fprintf(w, "route: " + route.Name)
             
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

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8000", nil)
}

func getMatchingRoute(path string, routes []RouteConfig) *RouteConfig {
    var matchingRoute *RouteConfig = nil
    for i, route := range routes {
        sampleRegex := regexp.MustCompile("^"+route.Path)
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
>>>>>>> 68f7fd6346f1a0fa3e020d287e27fc2341237346
