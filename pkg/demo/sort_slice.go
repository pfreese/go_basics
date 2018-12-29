package demo

import "sort"

// IntInSlice returns a boolean as to whether an integer is in a slice.
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// SortedUniqueIntSlice returns an ascending sorted slice of unique integers
// contained within an input slice.
func SortedUniqueIntSlice(s []int) []int {
	uniqueS := []int{}
	for _, val := range s {
		if !IntInSlice(val, uniqueS) {
			uniqueS = append(uniqueS, val)
		}
	}
	sort.Ints(uniqueS)
	return uniqueS
}