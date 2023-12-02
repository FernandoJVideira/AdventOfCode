package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	games := 0

	cubeColors := make(map[rune]int)
	for scanner.Scan() {

		var game int
		splitString := strings.Split(scanner.Text(), ":")
		fmt.Sscanf(splitString[0], "Game %d", &game)

		sets := strings.Split(splitString[1], ";")
		cubeColors['r'] = 0
		cubeColors['g'] = 0
		cubeColors['b'] = 0

		for _, set := range sets {
			cubes := strings.Split(set, ",")

			for _, cube := range cubes {

				var num int
				var color string

				_, err := fmt.Sscanf(cube, "%d %s", &num, &color)
				if err != nil {
					println(err)
				}
				if num > cubeColors[rune(color[0])] {
					cubeColors[rune(color[0])] = num
				}
			}
		}
		games += cubeColors['r'] * cubeColors['g'] * cubeColors['b']
	}
	fmt.Printf("Sum: %d\n", games)
}
