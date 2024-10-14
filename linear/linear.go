package linear

import "math"

//function to find the c in y=mx+c
func Intercept(data []float64) float64 {
	n := float64(len(data))
	var sumX, sumY, sumXY, sumX2, sumY2 float64
	for i, y := range data {
		x := float64(i)
		sumX += x // sum of x values
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y // sum of y^2 values
	}
	// Calculate the slope (m) and intercept (b) for the linear regression line
	// Formula: m = (n*Σ(x*y) - Σx*Σy) / (n*Σ(x^2) - (Σx)^2)
	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n
	return intercept
}

func Range(current float64, numbers []float64) (float64, float64) {
	m := 2.95

	var lower, upper float64

	c := Intercept(numbers)

	if current < c {
		lower = math.Max(0, c/m) - 15
		upper = c*m + 15
	} else {
		lower = c / m -15
		upper = c * m +15
	}

	return lower, upper
}
