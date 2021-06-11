package main

import (
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
