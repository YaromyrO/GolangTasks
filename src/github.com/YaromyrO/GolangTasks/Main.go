package main

import (
	"fmt"
	"./sorting/counting_sort"
	"math/rand"
	"time"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	slice := make([]int, 50)
		for i := range slice {
		slice[i] = r1.Intn(100)
	}

	fmt.Println(slice)
	fmt.Println(counting_sort.CountingSort(slice))

}
