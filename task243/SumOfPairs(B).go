package main

import "fmt"

func main() {

	fmt.Println(getPairs(200))
}

func getPairs(number int) []int {
	var result []int
	for i := 1; i * i <= number; i++ {
		for j := 1; j * j <= number; j++ {
			if i * i + j * j == number && i >= j{
				result = append(result, i, j)
			}
		}
	}
	return result
}



