package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntInSlice(t *testing.T) {
	// Test examples
	tables := []struct {
		val	int
		slice []int
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
		actual := IntInSlice(table.val, table.slice)
		if actual != table.expInList {
			t.Errorf("Error in determining if %v in %v", table.val, table.slice)
		}
	}
}

func TestSortedUniqueIntSlice(t *testing.T) {
	// Test examples
	tables := []struct {
		slice []int
		expSortedUnique	[]int
	}{
		// return original list if already sorted & unique
		{
			[]int{-2, 0, 3, 5, 100},
			[]int{-2, 0, 3, 5, 100},
		},
		// test repeated elements
		{
			[]int{-2, 0, 3, 5, 100, -2, 5, 5, 3, 100},
			[]int{-2, 0, 3, 5, 100},
		},
		// test repeated, unsorted elements
		{
			[]int{100, 5, -2, 5, 5, 3, -2, 100, 0},
			[]int{-2, 0, 3, 5, 100},
		},
	}

	// go through each test example
	for _, table := range tables {
		actual := SortedUniqueIntSlice(table.slice)
		if !reflect.DeepEqual(actual, table.expSortedUnique) {
			t.Errorf("Error in getting sorted unique slice: Wanted: %v, got: %v",
				table.expSortedUnique, actual)
		}
	}
}

func ExampleSortedUniqueIntSlice() {
	fmt.Println(SortedUniqueIntSlice([]int{100, 5, -2, 5, 5, 3, -2, 100, 0}))
	fmt.Println(SortedUniqueIntSlice([]int{-2, 0}))
	// Output: [-2 0 3 5 100]
	// [-2 0]
}