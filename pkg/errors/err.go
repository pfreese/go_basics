package errs

import (
	"fmt"
	"math"
)

// circleArea formats an error according to a format specifier and returns
// a string as value that satisfies error (i.e., implements the Error() method)
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// An alternative generic error is:
		// errors.New("area calculation failed, radius is <0")
		return 0, fmt.Errorf("area calculation failed, radius %0.2f is less than 0", radius)
	}
	area := math.Pi * radius * radius
	return area, nil
}
