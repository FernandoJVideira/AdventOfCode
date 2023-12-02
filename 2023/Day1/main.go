package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile(`\d`)
		line := re.FindAllString(scanner.Text(), -1)

		if len(line) == 1 {
			value, _ := strconv.Atoi(line[0] + line[0])
			sum += value
		} else {
			firstVal := line[0]
			lastVal := line[len(line)-1]

			result, _ := strconv.Atoi(firstVal + lastVal)
			sum += result
		}
	}

	fmt.Println(sum)
}
