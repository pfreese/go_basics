package errs

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	var rad float64 = 1
	actArea, actErr := circleArea(rad)

	var expectedErr error = nil
	if actErr != nil {
		t.Errorf("Error was incorrect, got: %v, want: %v.", actErr, expectedErr)
	}

	var expectedArea float64 = math.Pi
	if actArea != expectedArea {
		t.Errorf("Area was incorrect, got: %v, want: %v.", actArea, expectedArea)
	}
}