package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Move struct {
	Direction string
	Steps     int
}

type Dial struct {
	CurrentPosition int
}

func ParseInput() []Move {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	res := make([]Move, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dir := string(line[0])
		steps, _ := strconv.Atoi(line[1:])

		res = append(res, Move{
			Direction: dir,
			Steps:     steps,
		})
	}
	return res
}

func (d *Dial) TurnDial(direction string, steps int) int {
	count := 0
	switch direction {
	case "R":
		count = (d.CurrentPosition + steps) / 100

		d.CurrentPosition = (d.CurrentPosition + steps) % 100
	case "L":
		if d.CurrentPosition > 0 && steps >= d.CurrentPosition {
			count = 1 + (steps-d.CurrentPosition)/100
		} else if d.CurrentPosition == 0 {
			count = steps / 100
		}
		d.CurrentPosition = (d.CurrentPosition - steps) % 100
		if d.CurrentPosition < 0 {
			d.CurrentPosition += 100
		}
	}
	return count
}
