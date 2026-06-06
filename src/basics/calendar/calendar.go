package calendar

import "log"

func main() {
	event := Event{}

	err := event.SetTitle("An event")
	if err != nil {
		log.Fatal(err)
	}
}
