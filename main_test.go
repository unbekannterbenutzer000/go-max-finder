package main

import (
	"math"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name       string
		size       int
		wantLength int
	}{
		{"negative_size", -5, 0},
		{"zero_size", 0, 0},
		{"one_element", 1, 1},
		{"several_elements", 32, 32},
		{"larger_slice", 1000, 1000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)
			if len(data) != test.wantLength {
				t.Fatalf("len(generateRandomElements(%d)) = %d, want %d", test.size, len(data), test.wantLength)
			}
			for i, value := range data {
				if value < 1 || value > 1_000_000 {
					t.Fatalf("data[%d] = %d, want a value from 1 to 1000000", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil_slice", nil, 0},
		{"empty_slice", []int{}, 0},
		{"single_positive", []int{7}, 7},
		{"single_zero", []int{0}, 0},
		{"single_negative", []int{-7}, -7},
		{"equal_values", []int{4, 4, 4}, 4},
		{"repeated_maximum", []int{2, 9, 3, 9}, 9},
		{"maximum_at_start", []int{9, 4, 2}, 9},
		{"maximum_in_middle", []int{1, 9, 4}, 9},
		{"maximum_at_end", []int{1, 2, 9}, 9},
		{"zero_values", []int{0, 0, 0}, 0},
		{"negative_values", []int{-9, -2, -5}, -2},
		{"mixed_values", []int{-5, 0, 8, -2}, 8},
		{"minimum_int", []int{math.MinInt}, math.MinInt},
		{"int_limits", []int{math.MinInt, 0, math.MaxInt}, math.MaxInt},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum(%v) = %d, want %d", test.data, got, test.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil_slice", nil, 0},
		{"empty_slice", []int{}, 0},
		{"single_positive", []int{7}, 7},
		{"single_zero", []int{0}, 0},
		{"single_negative", []int{-7}, -7},
		{"equal_values", []int{4, 4, 4}, 4},
		{"repeated_maximum", []int{2, 9, 3, 9}, 9},
		{"maximum_at_start", []int{9, 4, 2}, 9},
		{"maximum_in_middle", []int{1, 9, 4}, 9},
		{"maximum_at_end", []int{1, 2, 9}, 9},
		{"zero_values", []int{0, 0, 0}, 0},
		{"negative_values", []int{-9, -2, -5}, -2},
		{"mixed_values", []int{-5, 0, 8, -2}, 8},
		{"seven_negative_values", []int{-9, -8, -7, -1, -6, -5, -4}, -1},
		{"exactly_eight_elements", []int{3, 8, 2, 1, 6, 5, 4, 7}, 8},
		{"nine_elements_maximum_in_remainder", []int{2, 3, 4, 5, 6, 7, 8, 1, 99}, 99},
		{"sixteen_elements", []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 16},
		{"seventeen_elements_maximum_in_remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 99}, 99},
		{"seventeen_negative_values", []int{-17, -16, -15, -14, -13, -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
		{"minimum_int", []int{math.MinInt}, math.MinInt},
		{"int_limits", []int{math.MinInt, 0, math.MaxInt}, math.MaxInt},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxChunks(test.data)
			if got != test.want {
				t.Fatalf("maxChunks(%v) = %d, want %d", test.data, got, test.want)
			}
			if sequentialMax := maximum(test.data); got != sequentialMax {
				t.Errorf("maxChunks(%v) = %d, maximum = %d", test.data, got, sequentialMax)
			}
		})
	}

	t.Run("maximum_in_each_position", func(t *testing.T) {
		for maxIndex := 0; maxIndex < 17; maxIndex++ {
			data := make([]int, 17)
			data[maxIndex] = 1
			if got := maxChunks(data); got != 1 {
				t.Fatalf("maxChunks with maximum at index %d = %d, want 1", maxIndex, got)
			}
		}
	})
}
