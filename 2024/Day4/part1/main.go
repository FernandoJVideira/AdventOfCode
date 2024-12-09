package main

import (
	"github.com/FernandoVideira/AdventOfCode/2024/Day4/part1/utils"
)

var word = "XMAS"
var directions = [][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}
var count = 0

func main() {
	grid := utils.ParseInput()

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				rowIndex := dir[0]
				colIndex := dir[1]
				isXmas := true

				for i := 0; i < len(word); i++ {
					rowOffset := row + (rowIndex * i)
					colOffset := col + (colIndex * i)

					if rowOffset < 0 || rowOffset >= len(grid) || colOffset < 0 || colOffset >= len(grid[row]) {
						isXmas = false
						break
					}

					if grid[rowOffset][colOffset] != rune(word[i]) {
						isXmas = false
						break
					}
				}
				if isXmas {
					count++
				}
			}
		}
	}

	println("XMAS Count:", count)
}
