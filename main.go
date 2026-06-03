package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var wg sync.WaitGroup

func main() {
	urls := os.Args[1:]
	workerCount := 5

	limiter := rate.NewLimiter(2, 5)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	// buffered channels
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	client := &http.Client{}

	// start workers
	for i := range workerCount {
		wg.Go(func() {
			worker(i, ctx, jobs, results, client, limiter)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for _, url := range urls {
			fmt.Printf("sending job %s\n", url)
			jobs <- Job{URL: url}
			fmt.Printf("sent job %s\n", url)
		}
		close(jobs)
	}()

	start := time.Now()

	// read results safely
	for res := range results {
		if res.Err != nil {
			fmt.Printf("[ERR] - %s\n", res.URL)
		} else {
			fmt.Printf("%s -> %d\n", res.URL, res.StatusCode)
		}
	}

	fmt.Printf("time: %d ms", time.Since(start).Milliseconds())
}
