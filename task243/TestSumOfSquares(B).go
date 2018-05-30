package main

import "testing"

func testSumOfSquaresB(t *testing.T) {
	var tests = []struct{
		input int
		expected []int
	}{
		{2450, []int{35, 35,49, 7}},
		{463, []int{}},
		{130, []int{9, 7, 11, 3}},
		{106, []int{9, 5}},
	}

	for _, test := range tests {
		if result := GetPairs(test.input); TestEq(result, test.expected) {
			t.Errorf("FindNum(%q) = %v", test.expected, result)
		}
	}
}

