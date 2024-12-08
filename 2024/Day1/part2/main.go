package main

import (
	"fmt"

	"github.com/FernandoVideira/AdventOfCode/2024/Day1/part2/utils"
)

func main() {

	var sum int

	left, right := utils.ParseInput()

	for _, v := range left {
		sum += utils.CountOcorrences(right, v) * v
	}

	fmt.Println(sum)
}
