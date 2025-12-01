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

func TestCountAllZeroPositions(t *testing.T) {
	dial := SafeDial{position: 50}

	rotations := []Rotation{
		{direction: Right, distance: 1000},
	}

	result := countAllZeroPositions(&dial, rotations)
	expected := 10

	if result != expected {
		t.Errorf("countAllZeroPositions() = %d, want %d", result, expected)
	}

	// Verify the dial ended at the correct position
	expectedFinalPos := 50
	if dial.position != expectedFinalPos {
		t.Errorf("Final position = %d, want %d", dial.position, expectedFinalPos)
	}
}

func TestCountAllZeroPositionsMultipleRotations(t *testing.T) {
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

	// Debug: track each rotation
	totalCount := 0
	for i, rot := range rotations {
		startPos := dial.position
		crossings := dial.move(rot.direction, rot.distance)
		totalCount += crossings
		t.Logf("Rotation %d: start=%d, dir=%v, dist=%d, end=%d, crossings=%d, total=%d",
			i, startPos, rot.direction, rot.distance, dial.position, crossings, totalCount)
	}

	result := countAllZeroPositions(&SafeDial{position: 50}, rotations)
	expected := 6

	if result != expected {
		t.Errorf("countAllZeroPositions() with multiple rotations = %d, want %d", result, expected)
	}
}

func TestCountAllZeroPositionsDetailed(t *testing.T) {
	tests := []struct {
		name          string
		startPos      int
		rotation      Rotation
		expectedCount int
		expectedEnd   int
	}{
		{
			name:          "position 50, R1000",
			startPos:      50,
			rotation:      Rotation{direction: Right, distance: 1000},
			expectedCount: 10,
			expectedEnd:   50,
		},
		{
			name:          "position 50, L68",
			startPos:      50,
			rotation:      Rotation{direction: Left, distance: 68},
			expectedCount: 1,
			expectedEnd:   82,
		},
		{
			name:          "position 82, L30",
			startPos:      82,
			rotation:      Rotation{direction: Left, distance: 30},
			expectedCount: 0,
			expectedEnd:   52,
		},
		{
			name:          "position 0, R50",
			startPos:      0,
			rotation:      Rotation{direction: Right, distance: 50},
			expectedCount: 0,
			expectedEnd:   50,
		},
		{
			name:          "position 0, R100",
			startPos:      0,
			rotation:      Rotation{direction: Right, distance: 100},
			expectedCount: 1,
			expectedEnd:   0,
		},
		{
			name:          "position 99, R1",
			startPos:      99,
			rotation:      Rotation{direction: Right, distance: 1},
			expectedCount: 1,
			expectedEnd:   0,
		},
		{
			name:          "position 1, L1",
			startPos:      1,
			rotation:      Rotation{direction: Left, distance: 1},
			expectedCount: 0,
			expectedEnd:   0,
		},
		{
			name:          "position 0, L1",
			startPos:      0,
			rotation:      Rotation{direction: Left, distance: 1},
			expectedCount: 0,
			expectedEnd:   99,
		},
		{
			name:          "position 5, L105",
			startPos:      5,
			rotation:      Rotation{direction: Left, distance: 105},
			expectedCount: 2,
			expectedEnd:   0,
		},
		{
			name:          "position 50, L50",
			startPos:      50,
			rotation:      Rotation{direction: Left, distance: 50},
			expectedCount: 0,
			expectedEnd:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dial := SafeDial{position: tt.startPos}
			count := dial.move(tt.rotation.direction, tt.rotation.distance)

			if count != tt.expectedCount {
				t.Errorf("Count = %d, want %d", count, tt.expectedCount)
			}
			if dial.position != tt.expectedEnd {
				t.Errorf("End position = %d, want %d", dial.position, tt.expectedEnd)
			}
		})
	}
}
