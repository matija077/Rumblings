package main

import (
	"net/http"

	"github.com/nsqio/go-nsq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func processEvent(message *nsq.Message) error {
	// Process the event here...
	opsProcessed.Inc()
	return nil
}

func main() {
	// Connect to NSQ
	consumer, err := nsq.NewConsumer("mytopic", "mychannel", nsq.NewConfig())
	if err != nil {
		// Handle error
		return
	}
	consumer.AddHandler(nsq.HandlerFunc(processEvent))
	err = consumer.ConnectToNSQD("localhost:4150")
	if err != nil {
		// Handle error
		return
	}

	defer consumer.DisconnectFromNSQD("localhost:4150")

	// Connect to StatsD
	/*statsdClient, err := statsd.NewClient("localhost:8125", "myapp.")
	if err != nil {
		// Handle error
		return
	}
	defer statsdClient.Close()*/

	// Create a Prometheus registry
	registry := prometheus.NewRegistry()
	// Register Prometheus metrics
	registry.MustRegister(opsProcessed)
	// Register StatsD metrics
	//registry.MustRegister(prometheus.NewStatsdCollector(statsdClient))

	// Expose Prometheus metrics at "/metrics"
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.ListenAndServe(":8080", nil)
}
