package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func DoRequest(f *FilterContext, w http.ResponseWriter, r *http.Request) {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}

	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest(r.Method, f.TargetURL+f.RequestURI, r.Body)
	req.Header = r.Header
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	} else {

		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}

		w.WriteHeader(resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		//Convert the body to type string
		sb := string(body)
		fmt.Fprint(w, sb)
	}

}
