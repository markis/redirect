package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	if config, ok := getConfig("config.json"); ok {
		http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
			host := strings.Split(req.Host, ":")[0]
			if redirect, ok := config.getRedirect(host); ok {
				var url = getPath(redirect.Host, req.URL.Path, req.URL.RawQuery)
				http.Redirect(w, req, url, redirect.Code|http.StatusMovedPermanently)
			} else {
				http.NotFound(w, req)
			}
		})
		port := config.Port
		log.Printf("About to listen on %v. Go to http://127.0.0.1:%v/", port, port)
		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getPath(host string, path string, query string) string {
	if len(query) > 0 {
		return host + path + "?" + query
	}
	return host + path
}
