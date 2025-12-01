# Advent of Code 2025

Solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Project Structure

```
aoc2025/
├── main.go              # Entry point to run solutions
├── days/
│   └── day01/
│       ├── solution.go      # Day 1 solution implementation
│       ├── solution_test.go # Day 1 tests
│       └── input.txt        # Day 1 puzzle input
└── pkg/
    └── utils/
        └── input.go         # Utility functions for reading input
```

## Usage

### Running a Solution

```bash
go run main.go <day>
```

Example:
```bash
go run main.go 1
```

### Running Tests

Run tests for a specific day:
```bash
go test ./days/day01/
```

Run all tests:
```bash
go test ./...
```

### Adding a New Day

1. Create a new directory: `mkdir -p days/dayXX`
2. Copy the template from day01 or create:
   - `solution.go` with `Part1()` and `Part2()` functions
   - `solution_test.go` with test cases
   - `input.txt` for your puzzle input
3. Add the new day case in `main.go`

## Development

- Go version: 1.21+ recommended
- No external dependencies required for basic solutions
