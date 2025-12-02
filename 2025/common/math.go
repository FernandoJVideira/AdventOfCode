package common

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MinOf returns the minimum value in a slice
func MinOf(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

// MaxOf returns the maximum value in a slice
func MaxOf(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

// Product returns the product of all integers in a slice
func Product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prod := 1
	for _, n := range nums {
		prod *= n
	}
	return prod
}

// GCD returns the greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Sum returns the sum of all integers in a slice
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// LCMOf returns the LCM of multiple numbers
func LCMOf(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for _, n := range nums[1:] {
		result = LCM(result, n)
	}
	return result
}

// Sign returns -1, 0, or 1 depending on the sign of x
func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// InRange checks if a value is within [min, max] inclusive
func InRange(val, minVal, maxVal int) bool {
	return val >= minVal && val <= maxVal
}

// Distance calculates the Manhattan distance between two points
func Distance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}
