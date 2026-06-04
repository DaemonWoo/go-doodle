package basics

import "fmt"

func formatter() {
	// %f - float
	fmt.Printf("A float: %f\n", 3.1415926535)

	// %7.3f - 7: minimal width; .3: max digits after `.`
	fmt.Printf("%7.3f\n", 1234.5678) // 1234.568

	// %d - decimal
	fmt.Printf("An integer: %d\n", 14)

	// %s - string
	fmt.Printf("A string; %s\n", "hello")

	// %t - logical
	fmt.Printf("A boolean: %t\n", true)

	// %v - variable(type is infered from passed value)
	fmt.Printf("Values: %v %v %v\n", 1, "\t", "true")

	// %#v - variable, but without formatted special symbols
	fmt.Printf("Values: %#v %#v %#v\n", "str", "\t", 100)

	// %T - type of the passed value
	fmt.Printf("Types: %T %T %T\n", "string", 147, false)

	// %% - % literal
	fmt.Printf("Percent sign: %%")
}
