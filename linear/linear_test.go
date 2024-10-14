package linear

import (
	"math"
	"testing"
)

// function to test the intercept ie c in y=mx+c
func TestIntercept(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	expectedIntercept := 1.0
	intercept := Intercept(data)

	if math.Abs(intercept-expectedIntercept) > 1e-9 {
		t.Errorf("Intercept(%v) = %v; want %v", data, intercept, expectedIntercept)
	}
}

func TestRange(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}

	// Test case when current is less than intercept
	current1 := 1.0
	lower1, upper1 := Range(current1, numbers)
	if lower1 >= upper1 {
		t.Errorf("Range(%v, %v) = (%v, %v); expected lower < upper", current1, numbers, lower1, upper1)
	}

	// Test case when current is greater than intercept
	current2 := 5.0
	lower2, upper2 := Range(current2, numbers)
	if lower2 >= upper2 {
		t.Errorf("Range(%v, %v) = (%v, %v); expected lower < upper", current2, numbers, lower2, upper2)
	}
}
