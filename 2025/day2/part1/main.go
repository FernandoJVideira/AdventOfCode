package main

import (
	"fmt"
	"path/filepath"

	"github.com/FernandoVideira/AdventOfCode/2025/common"
	"github.com/FernandoVideira/AdventOfCode/2025/day2"
)

func main() {
	inputPath := filepath.Join("..", "input.txt")
	lines := common.MustReadLinesSplit(inputPath, ",")

	res := day2.ParseInput(lines)

	sum := day2.SumInvalidIdsPart1(res)
	fmt.Println("Part 1 - Sum of invalid IDs is:", sum)
}
