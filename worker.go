package main

import "net/http"

func worker(jobs <-chan Job, results chan<- Result, client *http.Client) {
	for job := range jobs {
		res, err := client.Get(job.URL)
		if err != nil {
			results <- Result{URL: job.URL, Err: err}
			continue
		}

		defer res.Body.Close()

		results <- Result{
			URL:        job.URL,
			StatusCode: res.StatusCode,
		}
	}
}
