package utils

import (
	"math"
	"strconv"
	"strings"
)

func ParseLine(line string) []int {
	var nums []int
	parts := strings.Split(line, " ")
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		nums = append(nums, num)
	}
	return nums
}

func CalculateSlopeSlice(nums []int) []int {
	var diff []int
	for i := 0; i < len(nums)-1; i++ {
		diff = append(diff, nums[i+1]-nums[i])
	}
	return diff
}

func IsValidSlope(diff []int) bool {
	isPositive := diff[0] > 0
	for _, d := range diff {
		if (isPositive && d <= 0) || (!isPositive && d >= 0) || (math.Abs(float64(d)) > 3) {
			return false
		}
	}
	return true
}
