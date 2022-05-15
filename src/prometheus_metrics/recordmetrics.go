package prometheus_metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func Ping(w http.ResponseWriter, req *http.Request) {
	PingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func Vec(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	param1 := query.Get("param1")
	param2 := query.Get("param2")
	CounterVector.WithLabelValues(param1, param2).Inc()
	fmt.Fprintf(w, fmt.Sprintf("%s %s", param1, param2))
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	PingCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	})

	CounterVector = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "counter_vector_test",
		Help: "Counter vector test",
	}, []string{"method1", "method2"})
)
