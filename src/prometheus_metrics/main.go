package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	RecordMetrics()
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/g", G)
	http.HandleFunc("/vec", Vec)
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("starting server...")
	stopCh := make(chan error)
	go func() {
		stopCh <- http.ListenAndServe(":2112", nil)
	}()
	err := <-stopCh
	if err != nil {
		panic("sever crashed: " + err.Error())
	}
}
