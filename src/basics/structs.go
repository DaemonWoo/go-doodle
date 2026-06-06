package basics

import "fmt"

type User struct {
	Email string
	Name  string
	Address
}

var alice User

func init() {
	alice.Name = "Alice"
	alice.Email = "al1ce@mail.com"
	alice.Address = Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}

	fmt.Println(alice.Email)
	fmt.Println(alice.City)
}
