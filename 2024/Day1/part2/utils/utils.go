package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseInput() (left, right []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, " ")

		firstVal, _ := strconv.Atoi(tmp[0])
		scndVal, _ := strconv.Atoi(tmp[3])

		left = append(left, firstVal)
		right = append(right, scndVal)
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func CountOcorrences(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}
