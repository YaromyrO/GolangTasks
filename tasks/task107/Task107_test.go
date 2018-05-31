package task107

import "testing"

func TestSumOfSquaresB(t *testing.T) {
	var tests = []struct{
		input int
		expected int
	}{
		{12, 1},
		{23, 2},
		{56, 2},
		{100, 3},
	}

	for _, test := range tests {
		if result := FindNum(test.input); result != test.expected {
			t.Errorf("FindNum(%q) = %v", test.expected, result)
		}
	}
}
