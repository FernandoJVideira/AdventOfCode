package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FernandoVideira/AdventOfCode/2024/Day2/part2/utils"
)

var counter int = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := utils.ParseLine(line)
		diff := utils.CalculateSlopeSlice(nums)
		valid, i := utils.IsValidSlope(diff)
		if valid {
			counter++
		} else {
			valid = utils.VerifyAfterRemoving(nums, i)
			if valid {
				counter++
			}
		}
	}
	fmt.Println(counter)

}
