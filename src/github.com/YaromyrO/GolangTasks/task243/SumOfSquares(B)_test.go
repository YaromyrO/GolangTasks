package task243

import "testing"

func TestSumOfSquaresB(t *testing.T) {
	var tests = []struct{
		input int
		expected []int
	}{
		{2450, []int{35, 35,49, 7}},
		{130, []int{9, 7, 11, 3}},
		{106, []int{9, 5}},
		{444, []int{}},
	}

	for _, test := range tests {
		if result := GetPairs(test.input); !TestEq(result, test.expected) {
			t.Error("GetPairs(", test.input,") =", result)
		}
	}
}