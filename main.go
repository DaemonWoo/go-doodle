package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func checkURL(url string) {
	defer wg.Done()
	res, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error occured while fetching %s\n", url)
			return
		}
		res.Body.Close()
		
		fmt.Printf("%s -> %d\n", url, res.StatusCode)
}

func main() {
    urls := os.Args[1:]
	start := time.Now()

    for _, url := range urls {
		wg.Add(1)
        go checkURL(url)
    }

	wg.Wait()
	fmt.Printf("%d ms", time.Since(start).Milliseconds())
}