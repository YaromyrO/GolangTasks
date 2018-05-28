package main

import (
	"fmt"
	"math"
)

func main() {

	getPair(85)
}

func getPair(number int) {

	appears := false
	for x := 1; x < number/2; x++ {
		_, div1 := math.Modf(math.Sqrt(float64(x)))
		_, div2 := math.Modf(math.Sqrt(float64(number - x)))
		if div1 == 0 && div2 == 0 {
			appears = true
			fmt.Println("Pair for", number, "is:", math.Sqrt(float64(x)), "and", math.Sqrt(float64(number-x)))
			return
		}
	}

	if !appears {
		fmt.Println("No sum of squares for", number)
	}
}

