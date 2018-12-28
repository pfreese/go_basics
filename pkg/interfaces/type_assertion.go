package interfaces

import "fmt"

// Based on: https://golangbot.com/interfaces-part-1/

// Type assertion is used extract the underlying value of the interface.
// i.(T) is the syntax which is used to get the underlying value of
// interface i whose concrete type is T.

func AssertInt(i interface{}) {
	s := i.(int) // get the underlying int value from i
	fmt.Println(s)
}

// Safe type assertion

// - If the concrete type of i is T:
// 		v will have the underlying value of i
// 		ok will be true.
// - If the concrete type of i is not T:
// 		ok will be false
// 		v will have the zero value of type T (and the program will not panic)
func AssertIntSafe(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

