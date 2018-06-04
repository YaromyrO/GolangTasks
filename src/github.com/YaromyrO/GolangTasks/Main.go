package main

import (
	"./sorting/bucket_sort"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	slice := make([]int, 20)
	for i := range slice {
		slice[i] = r1.Intn(20)
	}
	fmt.Println(slice)
	//fmt.Println(quicksort.QuickSort(slice, quicksort.GetPivot))
	//fmt.Println(counting_sort.CountingSort(slice))
	//fmt.Println(merge_sort.MergeSort(slice))
	fmt.Println(bucket_sort.BucketSort(slice, bucket_sort.MsBits))

}
