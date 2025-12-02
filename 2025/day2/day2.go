package day2

import (
	"math"
	"strconv"
	"strings"

	"github.com/FernandoVideira/AdventOfCode/2025/common"
)

func ParseInput(lines [][]string) (res [][]int) {
	for _, line := range lines {
		for _, val := range line {
			nums := strings.Split(val, "-")
			start := common.ParseInt(nums[0])
			end := common.ParseInt(nums[1])
			res = append(res, []int{start, end})
		}
	}
	return res
}

func GenerateInvalidIDSumPart1(start, end, length int) int {
	if length%2 != 0 {
		return 0
	}
	sum := 0
	invalidID := make(map[int]bool)
	minPat := int(math.Pow10((length / 2) - 1))
	maxPat := int(math.Pow10(length/2)) - 1

	for p := minPat; p <= maxPat; p++ {
		patStr := strconv.Itoa(p)
		idStr := patStr + patStr
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		if id >= start && id <= end {
			if invalidID[id] {
				continue
			}
			sum += id
		}
		invalidID[id] = true

	}
	return sum

}

func GenerateInvalidIDSumPart2(start, end, length int) int {
	sum := 0
	invalidID := make(map[int]bool)
	for i := 1; i <= length/2; i++ {
		if length%i != 0 {
			continue
		}

		minPat := int(math.Pow10(i - 1))
		maxPat := int(math.Pow10(i)) - 1

		for p := minPat; p <= maxPat; p++ {
			patStr := strconv.Itoa(p)
			idStr := strings.Repeat(patStr, length/i)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				continue
			}
			if id >= start && id <= end {
				if invalidID[id] {
					continue
				}
				sum += id
			}
			invalidID[id] = true
		}
	}
	return sum
}

func SumInvalidIdsPart1(ranges [][]int) int {
	sum := 0
	for _, line := range ranges {
		start, end := line[0], line[1]

		sLen := len(strconv.Itoa(start))
		eLen := len(strconv.Itoa(end))

		for length := sLen; length <= eLen; length++ {
			sum += GenerateInvalidIDSumPart1(start, end, length)
		}
	}
	return sum
}

func SumInvalidIdsPart2(ranges [][]int) int {
	sum := 0
	for _, line := range ranges {
		start, end := line[0], line[1]

		sLen := len(strconv.Itoa(start))
		eLen := len(strconv.Itoa(end))

		for length := sLen; length <= eLen; length++ {
			sum += GenerateInvalidIDSumPart2(start, end, length)
		}
	}
	return sum
}
