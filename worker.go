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
) {
	for {
		select {
		case <-ctx.Done():
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

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
