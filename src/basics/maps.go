package basics

import "fmt"

func mapsSyntax() {
	ranks := map[string]int{
		"Meggy":  12,
		"Thomas": 11,
	}
	ranks["Harry"] = 30

	_, ok := ranks["Bob"]
	println(ok) // false

	for name, yo := range ranks {
		fmt.Printf("%s is %d years old\n", name, yo)
	}

	delete(ranks, "Harry")

	println(ranks)
}
