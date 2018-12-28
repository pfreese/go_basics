package errs

import (
	"errors"
	"math"
)

// circleArea formats an error according to a format specifier and returns
// a string as value that satisfies error (i.e., implements the Error() method)
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("area calculation failed, radius is <0")
	}
	area := math.Pi * radius * radius
	return area, nil
}
