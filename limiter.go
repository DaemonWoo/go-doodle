package main

import "time"

func Limiter() <-chan struct{} {
	// amount of requests per second
	const Rps = 2
	// token bucket capacity.
	// Allows up to {Burst} requests to be executed after idle time
	const Burst = 5

	limiter := make(chan struct{}, Burst)

	for range Burst {
		limiter <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(Rps))
		defer ticker.Stop()

		for range ticker.C {
			select {
			case limiter <- struct{}{}:
			default:
				// skip when buffer is full
			}
		}
	}()

	return limiter
}
