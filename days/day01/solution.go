// Package day01 contains the solution for Advent of Code 2025 Day 1.
package day01

import (
	"fmt"
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

func (s *SafeDial) move(direction Direction, distance int) {
	if direction == Left {
		s.position -= distance
	} else {
		s.position += distance
	}
	// Wrap around 0-99 range
	s.position = ((s.position % 100) + 100) % 100
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

func Part1(input string) int {
	rotations := parseRotations(input)
	dial := SafeDial{position: 50}
	return countZeroPositions(&dial, rotations)
}

func Part2(input string) int {
	// TODO: Implement Part 2 solution
	fmt.Println("Day 1, Part 2")
	return 0
}
