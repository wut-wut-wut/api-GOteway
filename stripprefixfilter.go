package main

import (
<<<<<<< HEAD
	"net/http"
	"strconv"
	"strings"
)

type StripPrefixFilter struct {
}

func (spf StripPrefixFilter) Process(f *FilterContext, w http.ResponseWriter, r *http.Request) error {
	depth, exists := f.Filter.Properties["depth"]
	d, _ := strconv.Atoi(depth)
	if exists {
		for cpt := 0; cpt <= d; cpt++ {
			i := strings.Index(f.RequestURI, `/`)
			if i > -1 {
				f.RequestURI = f.RequestURI[i+1:]
			}
		}

		if len(f.RequestURI) > 0 {
			if f.RequestURI[0] != '/' {
				f.RequestURI = `/` + f.RequestURI
			}
		} else {
			f.RequestURI = `/`
		}
	}
	return nil
}
=======
    "fmt"
    "strings"
    "strconv"
)

type StripPrefixFilter struct {

}

func (spf StripPrefixFilter)  Process(f *FilterContext) bool {
    fmt.Println("Process StripPrefixFilter")
    depth, exists := f.Filter.Properties["depth"]
     d, _ := strconv.Atoi(depth)
    if exists {
        fmt.Println("Depth : " + depth)
        fmt.Println("URI: " + f.RequestURI)
        for cpt := 0; cpt <= d; cpt++ {
            i := strings.Index(f.RequestURI, `/`)
            if i > -1 {
                fmt.Println("Index: ", i)
                f.RequestURI = f.RequestURI[i+1:]
            } else {
                fmt.Println("Index not found")
            }
        }
        
        if f.RequestURI[0] != '/' {
            f.RequestURI = `/` + f.RequestURI
        }
        
        fmt.Println(f.RequestURI)
    }
    return true
}
>>>>>>> 68f7fd6346f1a0fa3e020d287e27fc2341237346
