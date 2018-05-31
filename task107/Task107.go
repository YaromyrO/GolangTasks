package main

import (
	"math"
)

func FindNum(k int) int{
	result := 0
	if k > 1 {
		for i := 1; i < k; i++{
			if math.Pow(4, float64(i)) < float64(k) {
				result = i
			}
		}
		return result
	}
	return 0
}
