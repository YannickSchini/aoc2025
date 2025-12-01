package day01

import (
	"testing"
)

const exampleInput = ``

func TestPart1(t *testing.T) {
	result := Part1(exampleInput)
	expected := 0
	if result != expected {
		t.Errorf("Part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(exampleInput)
	expected := 0
	if result != expected {
		t.Errorf("Part2() = %d, want %d", result, expected)
	}
}

func TestSafeDialMove(t *testing.T) {
	tests := []struct {
		name            string
		initialPosition int
		direction       Direction
		distance        int
		expectedResult  int
	}{
		{
			name:            "position 11, rotate R8",
			initialPosition: 11,
			direction:       Right,
			distance:        8,
			expectedResult:  19,
		},
		{
			name:            "position 19, rotate L19",
			initialPosition: 19,
			direction:       Left,
			distance:        19,
			expectedResult:  0,
		},
		{
			name:            "position 5, rotate L10",
			initialPosition: 5,
			direction:       Left,
			distance:        10,
			expectedResult:  95,
		},
		{
			name:            "position 95, rotate R5",
			initialPosition: 95,
			direction:       Right,
			distance:        5,
			expectedResult:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dial := SafeDial{position: tt.initialPosition}
			dial.move(tt.direction, tt.distance)
			if dial.position != tt.expectedResult {
				t.Errorf("After move: got position %d, want %d", dial.position, tt.expectedResult)
			}
		})
	}
}

func TestSafeDialMultipleRotations(t *testing.T) {
	dial := SafeDial{position: 50}

	rotations := []Rotation{
		{direction: Left, distance: 68},
		{direction: Left, distance: 30},
		{direction: Right, distance: 48},
		{direction: Left, distance: 5},
		{direction: Right, distance: 60},
		{direction: Left, distance: 55},
		{direction: Left, distance: 1},
		{direction: Left, distance: 99},
		{direction: Right, distance: 14},
		{direction: Left, distance: 82},
	}

	for _, rot := range rotations {
		dial.move(rot.direction, rot.distance)
	}

	expected := 32
	if dial.position != expected {
		t.Errorf("After multiple rotations: got position %d, want %d", dial.position, expected)
	}
}

func TestCountZeroPositions(t *testing.T) {
	dial := SafeDial{position: 50}

	rotations := []Rotation{
		{direction: Left, distance: 68},
		{direction: Left, distance: 30},
		{direction: Right, distance: 48},
		{direction: Left, distance: 5},
		{direction: Right, distance: 60},
		{direction: Left, distance: 55},
		{direction: Left, distance: 1},
		{direction: Left, distance: 99},
		{direction: Right, distance: 14},
		{direction: Left, distance: 82},
	}

	result := countZeroPositions(&dial, rotations)
	expected := 3

	if result != expected {
		t.Errorf("countZeroPositions() = %d, want %d", result, expected)
	}
}

func TestParseRotations(t *testing.T) {
	input := `L29
R49
L41`

	rotations := parseRotations(input)

	if len(rotations) != 3 {
		t.Fatalf("Expected 3 rotations, got %d", len(rotations))
	}

	expected := []Rotation{
		{direction: Left, distance: 29},
		{direction: Right, distance: 49},
		{direction: Left, distance: 41},
	}

	for i, rot := range rotations {
		if rot.direction != expected[i].direction || rot.distance != expected[i].distance {
			t.Errorf("Rotation %d: got {%v, %d}, want {%v, %d}",
				i, rot.direction, rot.distance, expected[i].direction, expected[i].distance)
		}
	}
}
