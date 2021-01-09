package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	// prometheus logger
	// "github.com/prometheus/client_golang/prometheus/promhttp"
	 metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	 "github.com/slok/go-http-metrics/middleware"
	 "github.com/slok/go-http-metrics/middleware/std"
	 "github.com/prometheus/client_golang/prometheus/promhttp"
)
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}
/*ARandomFunction go a random function
*/
func ARandomFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fuck you")
 }
func main() {
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", ARandomFunction)
	h := std.Handler("", mdlw, mux)
	go func() {
		http.ListenAndServe(":9090", promhttp.Handler())
	}()
	http.ListenAndServe(":8080", h)
}