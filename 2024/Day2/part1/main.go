package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FernandoVideira/AdventOfCode/2024/Day2/part1/utils"
)

func main() {

	counter := 0

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
		fmt.Println(diff)

		if utils.IsValidSlope(diff) {
			fmt.Println("Valid")
			counter++
		}
	}

	fmt.Println(counter)

}
