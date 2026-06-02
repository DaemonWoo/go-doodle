package main

import "time"

func Limiter(rps int) <-chan struct{} {
	limiter := make(chan struct{}, rps)

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rps))
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