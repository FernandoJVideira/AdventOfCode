package day1

import (
	"github.com/FernandoVideira/AdventOfCode/2025/common"
)

type Move struct {
	Direction string
	Steps     int
}

type Dial struct {
	CurrentPosition int
}

func ParseInput(lines []string) []Move {
	res := make([]Move, 0, len(lines))
	for _, line := range lines {
		dir := string(line[0])
		steps := common.ParseInt(line[1:])
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

// TurnDialCountZeros rotates the dial and returns how many times it passed through 0
func (d *Dial) TurnDialCountZeros(direction string, steps int) int {
	count := 0
	switch direction {
	case "R":
		// Count how many times we cross 0 going right
		count = (d.CurrentPosition + steps) / 100
		d.CurrentPosition = (d.CurrentPosition + steps) % 100
	case "L":
		// Count how many times we cross 0 going left
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
