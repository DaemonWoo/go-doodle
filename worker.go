package main

import (
	"context"
	"fmt"
	"net/http"
)

func worker(
	id int,
	ctx context.Context,
	jobs <-chan Job,
	results chan<- Result,
	client *http.Client,
	limiter <-chan struct{},
) {
	for {
		select {
		case <-ctx.Done():
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			// Uncomment for rate-limiting logs
			//fmt.Printf("%s waiting token\n", job.URL)
			
			// rate limit (ctx-aware)
			select {
			case <-ctx.Done():
				return
			case <-limiter:
			}
			// Uncomment for rate-limiting logs
			// fmt.Printf("%s got token %s\n",
			// 	job.URL,
			// 	time.Now().Format("15:04:05.000"))

			fmt.Printf(
				"[worker-%d] processing %s\n",
				id,
				job.URL,
			)

			req, err := http.NewRequestWithContext(
				ctx,
				http.MethodGet,
				job.URL,
				nil,
			)
			if err != nil {
				results <- Result{
					URL: job.URL,
					Err: err,
				}
				continue
			}

			res, err := client.Do(req)

			fmt.Printf(
				"[worker-%d] finished %s\n",
				id,
				job.URL,
			)

			if err != nil {
				results <- Result{
					URL: job.URL,
					Err: err,
				}
				continue
			}
			res.Body.Close()

			results <- Result{
				URL:        job.URL,
				StatusCode: res.StatusCode,
			}
		}
	}
}
