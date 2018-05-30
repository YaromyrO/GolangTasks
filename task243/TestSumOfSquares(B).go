package main

import "testing"

func testSumOfSquaresB(t *testing.T) {
	expected := []int{10, 10, 14, 2}
	if !TestEq(expected, GetPairs(200)) {
		t.Error("GetPairs(200) = expected")
	}
}

