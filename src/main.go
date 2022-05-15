package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"testProject/src/prometheus_metrics"
)

func main() {
	println("Hello, world.")
	prometheus_metrics.RecordMetrics()
	http.HandleFunc("/ping", prometheus_metrics.Ping)
	http.HandleFunc("/vec", prometheus_metrics.Vec)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
