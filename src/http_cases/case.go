package main

import (
	"fmt"
	"net/http"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	http.HandleFunc("/tmpl", tmpl)
	fmt.Println("starting server...")
	err := server.ListenAndServe()
	// or
	// err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
