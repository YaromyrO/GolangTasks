package main

import "fmt"

func main() {

	fmt.Println(getPair(85))
}

func getPair(number int) (int, int) {

	for i := 1; i * i <= number; i++ {
		for j := 1; j * j <= number; j++ {
			if i * i + j * j == number{
				return i, j
			}
		}
	}
	return 0, 0
}

