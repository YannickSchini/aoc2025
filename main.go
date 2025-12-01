package main

import (
	"fmt"
	"log"
	"os"

	"aoc2025/days/day01"
	"aoc2025/pkg/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		fmt.Println("Example: go run main.go 1")
		os.Exit(1)
	}

	day := os.Args[1]

	switch day {
	case "1":
		runDay01()
	default:
		fmt.Printf("Day %s not implemented yet\n", day)
		os.Exit(1)
	}
}

func runDay01() {
	input, err := utils.ReadInput("days/day01/input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	fmt.Println("=== Day 1 ===")
	fmt.Printf("Part 1: %d\n", day01.Part1(input))
	fmt.Printf("Part 2: %d\n", day01.Part2(input))
}
