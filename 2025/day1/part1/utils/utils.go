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
	switch direction {
	case "R":
		// Move right (positive direction) on the dial with wrap-around
		d.CurrentPosition = (d.CurrentPosition + steps) % 100
	case "L":
		// Move left (negative direction) on the dial with wrap-around
		d.CurrentPosition = (d.CurrentPosition - steps) % 100
		if d.CurrentPosition < 0 {
			d.CurrentPosition += 100
		}
	}
	return d.CurrentPosition
}
