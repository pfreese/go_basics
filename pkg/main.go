package main

import "github.com/pfreese/go_basics/pkg/interfaces"

func main() {
	// Test functions in internal_representation.go
	var m interfaces.MyFloat = 4
	m.Test() // 4
	interfaces.DescribeTester(m) // "Interface type interfaces.MyFloat value 4"

	// Describe a generic interface
	s := "Hello World"
	interfaces.DescribeInterface(s) // "Type = string, value = Hello World"

	i := 55
	interfaces.DescribeInterface(i) // "Type = int, value = 55"

	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	interfaces.DescribeInterface(strt) // "Type = struct { name string }, value = {Naveen R}"

	//
	// Type assertions, in type_assertion.go
	//
	var b interface{} = 56
	interfaces.AssertInt(b) // prints 56 since b is an integer

	b = "Steven Paul"
	// interfaces.AssertInt(b) // panics since b is a string, not an integer

	interfaces.AssertIntSafe(b) // prints: "0 false"

	//
	// Type switches, in type_switch.go
	//
	interfaces.FindType("Naveen") // "I am a string and my value is Naveen"
	interfaces.FindType(77) // "I am an int and my value is 77"
	interfaces.FindType(89.98) // "Unknown type"
}
