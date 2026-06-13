package basics

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for range delay {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving roroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
