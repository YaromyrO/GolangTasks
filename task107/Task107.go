package main

import (
	"math"
	"fmt"
)

func main() {

	findNum(756)
}

func findNum(k int){
	result := 0
	if k > 1 {
		for i := 1; i < k; i++{
			if math.Pow(4, float64(i)) < float64(k) {
				result = i
			}
		}
		fmt.Println("Greates number is:", result)
	}
}
