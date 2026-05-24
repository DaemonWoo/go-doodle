package main

type Job struct {
	URL string
}

type Result struct {
	URL        string
	StatusCode int
	Err        error
}
