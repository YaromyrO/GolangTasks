package main

import "testing"

func testSumOfSquaresA(t *testing.T) {
	var tests = []struct{
		input int
		expected []int
	}{
		{32, []int{4, 4}},
		{74, []int{5, 7}},
		{40, []int{2, 6}},
		{85, []int{6, 7}},
		{444, []int{}},

	}

	for _, test := range tests {
		if result := GetPair(test.input); TestEq(result, test.expected) {
			t.Errorf("FindNum(%q) = %v", test.expected, result)
		}
	}
}
