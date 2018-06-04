package main

import (
	"fmt"
	"math/rand"
	"time"
	"./sorting/quicksort"
	"./sorting/merge_sort"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	slice := make([]int, 40)
	for i := range slice {
		slice[i] = r1.Intn(100)
	}

	fmt.Println(slice)
	fmt.Println(quicksort.QuickSort(slice, nil))
	//fmt.Println(counting_sort.CountingSort(slice))
	fmt.Println(merge_sort.MergeSort(slice))

}
