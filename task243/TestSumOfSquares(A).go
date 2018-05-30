package main

import "testing"

func testSumOfSquaresA(t *testing.T) {
	expected := []int{2, 9}
	if !TestEq(expected, GetPair(85)) {
		t.Error("GetPair(85) = expected")
	}
}

