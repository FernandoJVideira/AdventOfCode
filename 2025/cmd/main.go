package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	day := flag.Int("day", 1, "Day to run (1-12)")
	part := flag.Int("part", 1, "Part to run (1 or 2)")
	flag.Parse()

	if *day < 1 || *day > 12 {
		fmt.Println("Day must be between 1 and 12")
		os.Exit(1)
	}

	if *part < 1 || *part > 2 {
		fmt.Println("Part must be 1 or 2")
		os.Exit(1)
	}

	// Get the directory where the cmd is located
	execDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	// Go up one level to the 2025 directory if we're in cmd/
	if filepath.Base(execDir) == "cmd" {
		execDir = filepath.Dir(execDir)
	}

	// Build the path to the solution
	dayDir := fmt.Sprintf("day%d", *day)
	partDir := fmt.Sprintf("part%d", *part)
	solutionPath := filepath.Join(execDir, dayDir, partDir)

	// Check if the solution exists
	mainFile := filepath.Join(solutionPath, "main.go")
	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		fmt.Printf("Solution not found: %s\n", mainFile)
		os.Exit(1)
	}

	// Run the solution
	fmt.Printf("🎄 Running Day %d, Part %d...\n", *day, *part)
	fmt.Println(strings.Repeat("=", 50))

	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = solutionPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running solution: %v\n", err)
		os.Exit(1)
	}

}
