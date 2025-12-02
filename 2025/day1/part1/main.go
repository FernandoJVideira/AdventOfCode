package main

import (
	"path/filepath"

	"github.com/FernandoVideira/AdventOfCode/2025/common"
	"github.com/FernandoVideira/AdventOfCode/2025/day1"
)

func main() {
	inputPath := filepath.Join("..", "input.txt")
	lines := common.MustReadLines(inputPath)

	moves := day1.ParseInput(lines)
	dial := day1.Dial{
		CurrentPosition: 50,
	}

	password := 0

	for _, move := range moves {
		newPosition := dial.TurnDial(move.Direction, move.Steps)
		if newPosition == 0 {
			password++
		}
	}

	println("Part 1 - Password is:", password)
}
