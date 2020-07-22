package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"log"
	"math/rand"
	"time"
)

var addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")

var (
	c = promauto.NewCounter(prometheus.CounterOpts{
		Name: "codely_app_sample_metric",
		Help: "Sample metric for Codely course",
	})

	h = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "codely_app_sample_histogram",
		Help: "Sample histogram for Codely course",
	})

	d = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "codely_app_sample_devices",
		Help: "Sample counter opts devices for Codely course"}, []string{"device"})
)

func main() {

	rand.Seed(time.Now().UnixNano())
	r := gin.New()

	p := ginprometheus.NewPrometheus("http")
	p.Use(r)

	r.GET("/", func(c *gin.Context) {

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		switch rand.Intn(6) {
		case 0: c.JSON(200, "Hello world!")
		case 1: c.JSON(404, "Not Found!")
		case 2: c.JSON(500, "Oops!")
		case 3: c.JSON(401, "Unauthorized!")
		case 4: c.JSON(403, "Forbidden!")
		case 5: c.JSON(408, "Timeout!")
		default:
			c.JSON(200, "Hello world!")
		}
	})

	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			h.Observe(float64(rand.Intn(100-0+1) + 0))
			d.With(prometheus.Labels{"device":"/dev/sda"}).Inc()
			c.Inc()
			fmt.Print(".")
			time.Sleep(1 * time.Second)
		}
	}()

	log.Fatal(r.Run(*addr))
}
