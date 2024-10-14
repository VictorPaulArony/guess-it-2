package main

import (
	"fmt"

	"guess-it-2/linear"
)

func main() {
	var numbers []float64

	for {
		var input float64
		_, err := fmt.Scanf("%f", &input)
		if err != nil {
			fmt.Println("INVALID INPUT", err)
			break
		}
		numbers = append(numbers, input)
		if len(numbers) > 2 {
			numbers = numbers[1:]
		}
		if len(numbers) == 2 {
			lower, upper := linear.Range(input, numbers)
			fmt.Printf("%.f %.f\n", lower, upper)
		}

	}
}
