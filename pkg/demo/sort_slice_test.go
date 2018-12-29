package demo

import (
	"gotest.tools/assert"
	"testing"
)

func TestIntInSlice(t *testing.T) {
	// Test examples
	tables := []struct {
		val	int
		list []int
		expInList	bool
	}{
		{
			3,
			[]int{3, 5, 100},
			true,
		},
		{
			4,
			[]int{3, 5, 100},
			false,
		},
		{
			5,
			[]int{3, 5, 100},
			true,
		},
	}

	// go through each test example
	for _, table := range tables {
		actual := IntInSlice(table.val, table.list)
		assert.Equal(t, actual, table.expInList)
	}
}