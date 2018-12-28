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
}
