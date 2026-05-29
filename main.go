package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
    urls := os.Args[1:]
	workerCount := 5

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	jobs := make(chan Job)
	results := make(chan Result)

	client := &http.Client{}

	// start workers
    for range workerCount {
		wg.Go(func() {
			worker(ctx, jobs, results, client)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for _, url := range urls {
			jobs <- Job{URL: url}
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