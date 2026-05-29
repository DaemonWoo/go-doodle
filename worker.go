package main

import (
	"context"
	"net/http"
)

func worker(ctx context.Context, jobs <-chan Job, results chan<- Result, client *http.Client) {
	for {
		select {
		case <-ctx.Done():
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}
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
			if err != nil {
				results <- Result{
					URL: job.URL,
					Err: err,
				}
				continue
			}
			res.Body.Close()

			results <- Result{
				URL: job.URL,
				StatusCode: res.StatusCode,
			}
		}
	}
}
