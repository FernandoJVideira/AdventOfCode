package main

import (
	"github.com/FernandoVideira/AdventOfCode/2025/Day5/part2/utils"
)

func main() {
	input := utils.ParseInput()
	dial := utils.Dial{
		CurrentPosition: 50,
	}
	password := 0
	for move := range input {
		zero := dial.TurnDial(input[move].Direction, input[move].Steps)
		password += zero
	}

	println(password)
}
