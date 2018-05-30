package main

import "testing"

func testSumOfSquaresB(t *testing.T) {
	expected := 4
	if expected == FindNum(756) {
		t.Error("FindNum(756) = expected")
	}
}
