package main

import (
	"fmt"
	//"go.uuid"
	//"grpc-go/benchmark/stats"
	"log"
	"mux"
	"net/http"
	//"ratecounter"
	"metrics/ratecounter"
	"metrics/track2"
	"time"
)

type metricsData struct {
	// TODO: on service
	// upstreamTimeMax
	// upstreamTimeAverage

	responseTimeTracker      *track2.Tracker
	requestCounter           *ratecounter.RateCounter
	requestConcurrentCounter *ratecounter.RateCounter
	requestConcurrentTracker *track2.Tracker
}

var data metricsData

//var counter *ratecounter.RateCounter

//var counter2 *stats.Counter
//var tracker *stats.Tracker

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		starttime := time.Now()

		fmt.Fprintf(w, "Hi there,  I love %s!\n", r.URL.Path[1:])
		data.requestCounter.Incr(1)
		fmt.Fprintf(w, "The number of http requests per minute for / is %d\n", data.requestCounter.Rate())

		//w.Header().Set("Content-Type", "text/html")
		// calls h(w,r)

		//h.ServeHTTP(w, r)

		responsetime := time.Since(starttime).Seconds() * 1000000
		microresponsetime := responsetime
		data.responseTimeTracker.Add(microresponsetime)
		fmt.Fprintf(w, "The current response time is %6.2f microseconds\n", microresponsetime)
		fmt.Fprintf(w, "The maximal response time for the last minute is %6.2f microseconds\n", data.responseTimeTracker.Max())
		fmt.Fprintf(w, "The average response time for the last minute is %6.2f microseconds\n", data.responseTimeTracker.Average())

	})

}

/**
*   Handler functions
*
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	counter.Incr(1)
	fmt.Fprintf(w, "Hi there,  I love %s!\n", r.URL.Path[1:])
	fmt.Fprintf(w, "The number of http requests per minute for / is %d\n", counter.Rate())
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there ,I love %s!", r.URL.Path[1:])
}
*/

func handlerICon(w http.ResponseWriter, r *http.Request) {
}
func main() {

	/*
	   Implementation counters and stats of 'requests-per-second' (for example)
	*/
	//	counter = ratecounter.NewRateCounter(60 * time.Second)
	//	counter2 = stats.NewCounter()
	//	counter2.Set()

	/*
	   Implementation of tracking maximum request time
	*/
	//	tracker = stats.NewTracker()

	/*
	   multiplexer dispatcher to redirect http requests
	*/

	data.requestCounter = ratecounter.NewRateCounter(60 * time.Second)
	data.requestConcurrentCounter = ratecounter.NewRateCounter(60 * time.Second)
	data.requestConcurrentTracker = track2.NewTracker(60 * time.Second)
	data.responseTimeTracker = track2.NewTracker(60 * time.Second)

	r := mux.NewRouter()
	/**
		  *   Used in case of handler functions
		  *
		  	r.HandleFunc("/", HomeHandler)
	                r.HandleFunc("/products", ProductsHandler)
	*/
	http.Handle("/", LoggingMiddleware(r))
	// Google chrome tries to find out favicon.ico, therefor if not handled the handler is called twice
	http.HandleFunc("/favicon.ico", handlerICon)
	/*
	   website server
	   if any buildin error occurs to ListenAndServe, it prints it and calls os.Exit(1)
	*/
	log.Fatal(http.ListenAndServe(":9876", nil))

}
