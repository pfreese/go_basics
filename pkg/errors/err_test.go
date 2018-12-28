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
		// Radius of -1; error
		{
			-1,
			errors.New("area calculation failed, radius is <0"),
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
		} else {
			// Check that the error message is as expected
			assert.ErrorContains(t, err, "area calculation failed, radius is <0")
		}
	}
}