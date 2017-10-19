package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

func main() {
	// Parse the command-line flags.
	flag.Parse()

	// Start the dispatcher.
	var wg sync.WaitGroup
	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers, &wg)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", Collector)

	// monitor signal 
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	go monitor(stopChan, wg)
	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}

func monitor(stopChan chan os.Signal, wg sync.WaitGroup) {
	fmt.Println("Waiting signal ...")
	<- stopChan
	StopDispatcher()
	fmt.Println("Server is shutting down ...")
	wg.Wait()
}
