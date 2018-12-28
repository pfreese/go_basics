package interfaces

import "fmt"

// Based on: https://golangbot.com/interfaces-part-1/

// Interface internal representation
// An interface can be thought of as being represented internally by a
// tuple (type, value).
// - type is the underlying concrete type of the interface
// - value holds the value of the concrete type.

type Tester interface {
	Test()
}

type MyFloat float64

// define a Test method on MyFloat type
func (m MyFloat) Test() {
	fmt.Println(m)
}

func DescribeTester(t Tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}


// The Empty Interface
// An interface which has zero methods is called empty interface.
// It is represented as interface{}.
// Since the empty interface has zero methods, all types implement the empty interface.

// DescribeInterface can be passed any type
func DescribeInterface(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

