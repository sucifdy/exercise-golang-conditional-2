package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var idealWeight float64

	if gender == "laki-laki" {
		idealWeight = float64(height-100) * 0.9 // 10% dari (height - 100)
	} else if gender == "perempuan" {
		idealWeight = float64(height-100) * 0.85 // 15% dari (height - 100)
	}

	return idealWeight
}

// debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 170)) // Output: 63
	fmt.Println(BMICalculator("perempuan", 165)) // Output: 55.25
}
