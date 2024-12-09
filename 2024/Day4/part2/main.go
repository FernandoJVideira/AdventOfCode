package main

import (
	"github.com/FernandoVideira/AdventOfCode/2024/Day4/part2/utils"
)

var directions = [][2][2]int{
	{
		{-1, 1},
		{1, -1},
	},
	{
		{-1, -1},
		{1, 1},
	},
}
var count = 0

func main() {
	grid := utils.ParseInput()

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != 'A' {
				continue
			}
			isXmas := true
			for _, dir := range directions {
				row1Index := row + dir[0][0]
				col1Index := col + dir[0][1]

				row2Index := row + dir[1][0]
				col2Index := col + dir[1][1]

				if row1Index < 0 || row1Index >= len(grid) || col1Index < 0 || col1Index >= len(grid[row]) || row2Index < 0 || row2Index >= len(grid) || col2Index < 0 || col2Index >= len(grid[row]) {
					isXmas = false
					break
				}

				if (grid[row1Index][col1Index] == 'M' && grid[row2Index][col2Index] == 'S') || (grid[row1Index][col1Index] == 'S' && grid[row2Index][col2Index] == 'M') {
					continue
				}

				isXmas = false
				break
			}
			if isXmas {
				count++
			}
		}
	}

	println("XMAS Count:", count)
}
