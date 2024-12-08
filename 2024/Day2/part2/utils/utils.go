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

func IsValidSlope(diff []int) (bool, int) {
	isPositive := diff[0] > 0
	for i, d := range diff {
		if (isPositive && d <= 0) || (!isPositive && d >= 0) || (math.Abs(float64(d)) > 3) {
			return false, i
		}
	}
	return true, 0
}

func VerifyAfterRemoving(nums []int, i int) bool {
	indicesToCheck := []int{i, i + 1, i - 1}
	for _, index := range indicesToCheck {
		if index >= 0 && index < len(nums) {
			tmp := Remove(nums, index)
			diff := CalculateSlopeSlice(tmp)
			valid, _ := IsValidSlope(diff)
			if valid {
				return true
			}
		}
	}
	return false
}

func Remove(slice []int, s int) []int {
	// Create a copy of the slice
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)

	// Verify if removing the last element
	if s == len(copySlice)-1 {
		return copySlice[:s]
	}

	if s == 0 {
		return copySlice[1:]
	}
	return append(copySlice[:s], copySlice[s+1:]...)
}
