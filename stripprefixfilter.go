package main

import (
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