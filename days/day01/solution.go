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

	// Calculate the raw new position (can be negative or >= 100)
	var rawNewPos int
	if direction == Right {
		rawNewPos = startPos + distance
	} else {
		rawNewPos = startPos - distance
	}

	// Count how many times we pass through a multiple of 100
	// Key insight: crossing from 99→0 or 0→99 means passing through a multiple of 100
	var boundaryCrossings int

	if direction == Right {
		// Moving right: simple case
		boundaryCrossings = rawNewPos/100 - startPos/100
	} else {
		// Moving left: need to handle negatives
		if rawNewPos >= 0 {
			// Still in positive range, no crossings
			boundaryCrossings = 0
		} else if startPos == 0 {
			// Starting from 0 and going negative
			// Only count full loops: -1 to -100 is 0 crossings, -101 to -200 is 1 crossing, etc.
			boundaryCrossings = (-rawNewPos - 1) / 100
		} else {
			// Starting from positive, going negative
			// We cross 0 once, then count additional full loops
			// Each additional 100 steps in negative direction crosses 0 again
			boundaryCrossings = 1 + (-rawNewPos) / 100
		}
	}

	// Wrap position to 0-99 range
	s.position = ((rawNewPos % 100) + 100) % 100

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
