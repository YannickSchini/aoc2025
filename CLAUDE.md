# Repository Overview

Advent of Code 2025 solutions written in Go. This repository uses **Jujutsu (jj)** for version control, not Git.

# Common Commands

## Running Solutions
```bash
go run main.go <day>        # Run a specific day's solution (e.g., go run main.go 1)
```

## Testing
```bash
go test ./days/day01/       # Run tests for a specific day
go test ./...               # Run all tests
go test -v ./days/day01/    # Run tests with verbose output
```

# Architecture

## Solution Structure
Each day follows the same pattern:
- Each day has its own package in `days/dayXX/`
- Solutions export two functions: `Part1(input string) int` and `Part2(input string) int`
- Both functions take the raw input as a string and return an integer result
- Test files use the example input from the problem statement

## Adding a New Day
1. Create directory: `days/dayXX/` (e.g., `days/day02/`)
2. Create three files:
   - `solution.go`: Implement `Part1()` and `Part2()` functions in package `dayXX`
   - `solution_test.go`: Add test cases with example input
   - `input.txt`: Paste puzzle input from adventofcode.com
3. Update `main.go`:
   - Import the new day package
   - Add case to switch statement
   - Create `runDayXX()` function that reads input and calls both parts

## Entry Point Flow
`main.go` → switch on day number → `runDayXX()` → reads `input.txt` → calls `Part1()` and `Part2()`

## Utilities
`pkg/utils/input.go` provides helper functions:
- `ReadInput(filepath)`: Returns entire file as string
- `ReadLines(filepath)`: Returns file as slice of lines
