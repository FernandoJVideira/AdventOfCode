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

	possibleGames := 0

	cubeColors := make(map[rune]int)
	for scanner.Scan() {

		var game int
		splitString := strings.Split(scanner.Text(), ":")
		fmt.Sscanf(splitString[0], "Game %d", &game)

		sets := strings.Split(splitString[1], ";")
		isValid := true

		for _, set := range sets {
			cubeColors['r'] = 0
			cubeColors['g'] = 0
			cubeColors['b'] = 0

			cubes := strings.Split(set, ",")

			for _, cube := range cubes {

				var num int
				var color string

				_, err := fmt.Sscanf(cube, "%d %s", &num, &color)
				if err != nil {
					println(err)
				}

				colorRune := rune(color[0])

				cubeColors[colorRune] += num

			}
			if cubeColors['r'] > 12 || cubeColors['g'] > 13 || cubeColors['b'] > 14 {
				isValid = false
				break
			}
		}

		if isValid {
			possibleGames += game
		}
	}
	fmt.Printf("Sum: %d\n", possibleGames)
}
