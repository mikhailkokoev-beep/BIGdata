package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		expectedLen int
	}{
		{"Negative size", -5, 0},
		{"Zero size", 0, 0},
		{"Small size", 10, 10},
		{"Large size", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			if len(result) != tt.expectedLen {
				t.Errorf("Длина слайса = %d, ожидалось %d", len(result), tt.expectedLen)
			}

			for i, val := range result {
				if val <= 0 {
					t.Errorf("Элемент %d = %d, должен быть > 0", i, val)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"Multiple elements", []int{1, 5, 3, 9, 2}, 9},
		{"All same", []int{7, 7, 7}, 7},
		{"Max at start", []int{100, 2, 3}, 100},
		{"Max at end", []int{1, 2, 100}, 100},
		{"Max in middle", []int{1, 100, 2}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Less than CHUNKS", []int{5, 2, 9}, 9},
		{"Exact CHUNKS", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"Multiple of CHUNKS", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 16},
		{"With remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"Max in first chunk", []int{100, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 100},
		{"Max in last chunk", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 100}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			if result != tt.expected {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.input, result, tt.expected)
			}
		})
	}
}