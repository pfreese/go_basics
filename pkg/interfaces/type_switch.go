package interfaces

import "fmt"

// Based on: https://golangbot.com/interfaces-part-1/

// A type switch is used to the compare the concrete type of an interface against multiple
// types specified in various case statements. It is similar to switch case. The only
// difference being the cases specify types and not values as in normal switch.

func FindType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

