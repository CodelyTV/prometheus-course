package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")

var (
	c = promauto.NewCounter(prometheus.CounterOpts{
		Name: "codely_app_sample_metric",
		Help: "Sample counter for Codely course",
	})

	h = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "codely_app_sample_histogram",
		Help: "Sample histogram for Codely course",
	})

	d = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "codely_app_sample_devices",
		Help: "Sample counter with devices label for Codely course"}, []string{"device"})

	e = promauto.NewCounter(prometheus.CounterOpts{
		Name: "codely_app_push_metric",
		Help: "Sample metric for Codely course (push)",
	})
)

func main() {

	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			h.Observe(float64(rand.Intn(100-0+1) + 0))
			d.With(prometheus.Labels{"device":"/dev/sda"}).Inc()
			c.Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			// Example of metric push
			_ = push.New("http://pushgateway:9091", "codely_job").Collector(e).Add()
			e.Inc()
			fmt.Print("_")
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
