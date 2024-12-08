package main

import (
	"fmt"
	"math"

	"github.com/FernandoVideira/AdventOfCode/2024/Day1/part1/utils"
)

func main() {
	var sum int
	left, right := utils.ParseInput()

	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println(sum)
}
