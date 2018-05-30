package main

import (
	"testing"
)

func TestSumOfSquaresA(t *testing.T) {
	var tests = []struct{
		input int
		expected []int
	}{
		{32, []int{4, 4}},
		{74, []int{5, 7}},
		{40, []int{2, 6}},
		{85, []int{2, 9}},
		{444, []int{}},

	}

	for _, test := range tests {
		if result := GetPair(test.input); !TestEq(result, test.expected) {
			t.Error("GetPair(", test.input,") =", result)
		}
	}
}