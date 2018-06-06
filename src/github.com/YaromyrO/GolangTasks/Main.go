package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/YaromyrO/GolangTasks/src/github.com/YaromyrO/GolangTasks/sorting/bucket_sort"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	slice := make([]int, 10000)
	for i := range slice {
		slice[i] = r1.Intn(1000)
	}
	//slice := []int{10, 15, 7, 5, 13, 22, 6, 9}
	fmt.Println(slice)
	//quicksort.QuickSort(slice, quicksort.GetPivot)
	//fmt.Println

	//fmt.Println(counting_sort.CountingSort(slice))
	//fmt.Println(merge_sort.MergeSort(slice))
	before := time.Now()
	fmt.Println(bucket_sort.BucketSort(slice, bucket_sort.MsBits))
	after := time.Now()
	println(after.Sub(before) /1000000)

}
