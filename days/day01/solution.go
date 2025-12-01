// Package day01 contains the solution for Advent of Code 2025 Day 1.
package day01

import (
	"strconv"
	"strings"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type SafeDial struct {
	position int
}

func (s *SafeDial) move(direction Direction, distance int) int {
	startPos := s.position

	// Update position first
	if direction == Left {
		s.position -= distance
	} else {
		s.position += distance
	}

	// Count how many times we crossed through 0 (boundary crossings, not including final position)
	// When wrapping, we cross boundaries. The number of full wraps is abs(position) / 100
	var boundaryCrossings int
	if direction == Right {
		// Moving right: count crossings from 99→0
		boundaryCrossings = (startPos + distance) / 100
		if startPos > 0 {
			boundaryCrossings -= startPos / 100
		}
	} else {
		// Moving left: count crossings through 0 (going backwards past 0)
		if distance > startPos {
			// We wrap around
			if startPos == 0 {
				// Starting from 0: only count if we make full loops back to 0
				boundaryCrossings = distance / 100
			} else {
				// Not starting from 0: count initial crossing plus full loops
				remaining := distance - startPos
				boundaryCrossings = 1 + remaining / 100
			}
		} else {
			// No wrapping
			boundaryCrossings = 0
		}
	}

	// Wrap around 0-99 range
	s.position = ((s.position % 100) + 100) % 100

	return boundaryCrossings
}

type Rotation struct {
	direction Direction
	distance  int
}

func parseRotations(input string) []Rotation {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rotations := make([]Rotation, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		dirChar := line[0]
		distStr := line[1:]
		distance, err := strconv.Atoi(distStr)
		if err != nil {
			continue
		}

		var direction Direction
		switch dirChar {
		case 'L':
			direction = Left
		case 'R':
			direction = Right
		default:
			continue
		}

		rotations = append(rotations, Rotation{direction: direction, distance: distance})
	}

	return rotations
}

func countZeroPositions(dial *SafeDial, rotations []Rotation) int {
	count := 0
	for _, rot := range rotations {
		dial.move(rot.direction, rot.distance)
		if dial.position == 0 {
			count++
		}
	}
	return count
}

func countAllZeroPositions(dial *SafeDial, rotations []Rotation) int {
	count := 0
	for _, rot := range rotations {
		startPos := dial.position

		// Count boundary crossings
		crossings := dial.move(rot.direction, rot.distance)
		count += crossings

		// Also count if we stopped at 0, BUT only if we didn't cross a boundary to get there
		// If we crossed a boundary and landed on 0, the crossing already counted it
		if dial.position == 0 {
			// Check if the crossing would have landed us on 0
			if crossings == 0 || (startPos == 0 && crossings > 0) {
				// Either no crossings (direct move to 0), or started from 0 and came back
				count++
			}
		}
	}
	return count
}

func Part1(input string) int {
	rotations := parseRotations(input)
	dial := SafeDial{position: 50}
	return countZeroPositions(&dial, rotations)
}

func Part2(input string) int {
	rotations := parseRotations(input)
	dial := SafeDial{position: 50}
	return countAllZeroPositions(&dial, rotations)
}
