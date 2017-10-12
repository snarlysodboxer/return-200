package main

import (
	"flag"
	"github.com/didip/tollbooth"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	var listenURL = flag.String("listen-url", "http://localhost:8080", "address to listen and serve upon")
	var rateLimit = flag.Int64("rate-limit", 1000, "milliseconds at which to limit the rate of requests")
	var burstLimit = flag.Int64("burst-limit", 10, "quantity of requests to allow in bursts")
	flag.Parse()

	parsedURL, err := url.Parse(*listenURL)
	if err != nil {
		log.Fatal(err)
	}
	path := "/"
	if parsedURL.Path != "" {
		path = parsedURL.Path
	}
	log.Printf("Listening on %s at %s", parsedURL.Host, path)
	log.Printf("Rate limiting to 1 request per %d milliseconds, with bursts of up to %d requests at once", *rateLimit, *burstLimit)

	mux := http.NewServeMux()
	mux.Handle(path,
		tollbooth.LimitFuncHandler(
			tollbooth.NewLimiter(*burstLimit, time.Duration(*rateLimit)*time.Millisecond),
			func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(200)
			},
		),
	)

	http.ListenAndServe(parsedURL.Host, mux)
}
