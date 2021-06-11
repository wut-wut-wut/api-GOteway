package main

import (
	"errors"
	"net/http"
)

type BasicAuthFilter struct {
}

func (baf BasicAuthFilter) Process(f *FilterContext, w http.ResponseWriter, r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return errors.New("no authorization header available")
	} else {
		if !isAuthorized(username, password, f) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return errors.New("bad credentials")
		}
	}

	return nil
}

func isAuthorized(u string, p string, f *FilterContext) bool {
	return f.Filter.Properties["username"] == u && f.Filter.Properties["password"] == p
}
