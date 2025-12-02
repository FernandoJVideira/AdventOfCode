package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads all lines from a file
func ReadLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

// MustReadLines reads lines or panics (useful for AoC)
func MustReadLines(filepath string) []string {
	lines, err := ReadLines(filepath)
	if err != nil {
		panic(err)
	}
	return lines
}

// ReadFile reads entire file as a string
func ReadFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// MustReadFile reads file or panics
func MustReadFile(filepath string) string {
	content, err := ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return content
}

// ReadGroups reads file and splits by blank lines
func ReadGroups(filepath string) ([][]string, error) {
	lines, err := ReadLines(filepath)
	if err != nil {
		return nil, err
	}

	var groups [][]string
	var currentGroup []string

	for _, line := range lines {
		if line == "" {
			if len(currentGroup) > 0 {
				groups = append(groups, currentGroup)
				currentGroup = nil
			}
		} else {
			currentGroup = append(currentGroup, line)
		}
	}

	if len(currentGroup) > 0 {
		groups = append(groups, currentGroup)
	}

	return groups, nil
}

// MustReadGroups reads groups or panics
func MustReadGroups(filepath string) [][]string {
	groups, err := ReadGroups(filepath)
	if err != nil {
		panic(err)
	}
	return groups
}

// ReadLinesSplit reads lines and splits each line by a separator
func ReadLinesSplit(filepath, separator string) ([][]string, error) {
	lines, err := ReadLines(filepath)
	if err != nil {
		return nil, err
	}

	result := make([][]string, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue // skip empty lines
		}
		parts := strings.Split(line, separator)
		result = append(result, parts)
	}

	return result, nil
}

// MustReadLinesSplit reads and splits lines or panics
func MustReadLinesSplit(filepath, separator string) [][]string {
	result, err := ReadLinesSplit(filepath, separator)
	if err != nil {
		panic(err)
	}
	return result
}

// ReadInts reads a file where each line is an integer
func ReadInts(filepath string) ([]int, error) {
	lines, err := ReadLines(filepath)
	if err != nil {
		return nil, err
	}

	ints := make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, nil
}

// MustReadInts reads integers or panics
func MustReadInts(filepath string) []int {
	ints, err := ReadInts(filepath)
	if err != nil {
		panic(err)
	}
	return ints
}

// ParseInt parses an integer, ignoring errors
func ParseInt(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}
