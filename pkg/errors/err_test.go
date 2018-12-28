package errs

import (
	"gotest.tools/assert"
	"errors"
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	// Test examples
	tables := []struct {
		r float64
		expErr	error
		expArea	float64
	}{
		// Radius of 1; no error
		{
			1,
			nil,
			math.Pi,
		},
		// Radius of 3; no error
		{
			3,
			nil,
			9 * math.Pi,
		},
		// Radius of 0; no error
		{
			0,
			nil,
			0,
		},
		// Radius of -1; error
		{
			-1,
			errors.New("area calculation failed, radius -1.00 is less than 0"),
			0,
		},
	}

	// Go through each test example
	for _, table := range tables {
		area, err := circleArea(table.r)
		// If no error expected, check that the area is within acceptable tolerance
		if table.expErr == nil {
			assert.NilError(t, err)
			assert.Assert(t, math.Abs(area - table.expArea) < 0.0001)
		// Otherwise check that the error message is as expected
		} else {
			assert.ErrorContains(t, err, table.expErr.Error())
		}
	}
}