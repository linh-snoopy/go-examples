package main

import "fmt"
import "sync"

var WorkerQueue chan chan WorkRequest
var workers []Worker

func StartDispatcher(nworkers int, wg *sync.WaitGroup) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start(wg)
		wg.Add(1)
		workers = append(workers, worker)
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}

func StopDispatcher() {
	for i, _ := range workers {
		workers[i].Stop()
	}
}
