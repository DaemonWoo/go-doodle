package basics 

func createPointer() *int64 {
	a := int64(5)
	return &a
}

func printBoolPointer(b *bool) {
	println(*b)
}

