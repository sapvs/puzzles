package prime

import "math"

// PrimeV1 worst
func PrimeV1(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// PrimeV2 better
func PrimeV2(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// PrimeV3 best
func PrimeV3(num int) bool {
	if num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num > 2 && num%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Floor(math.Sqrt(float64(num)))); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
