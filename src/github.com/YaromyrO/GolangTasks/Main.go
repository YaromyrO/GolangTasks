package main

import (
	"./sorting/counting_sort"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	slice := make([]int, 50)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100 - 1) + 1
	}
	fmt.Println(slice)
	fmt.Println(counting_sort.CountingSort(slice))
}
