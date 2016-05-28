// Copyright 2016 By GuanLuyong.  All rights reserved.

// Package main is entry of Web server.
package main

import (
	"com/guanly/lab/gweb/router"
	"fmt"
	"net/http"
)

const (
	Host = "0.0.0.0"
	Port = "8080"
)

func main() {
	srvMux := http.NewServeMux()

	// static server
	srvMux.Handle("/web", http.FileServer(http.Dir("webview")))
	srvMux.Handle("/wap", http.FileServer(http.Dir("wapview")))

	// register routers
	srvMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go! \n%s", "First Go Programe.")
	})
	// Login router
	router.RegisterUtil(router.LoginRouter{srvMux})

	http.ListenAndServe(fmt.Sprintf("%s:%s", Host, Port), srvMux)
}
