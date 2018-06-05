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

	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = r1.Intn(200)
	}
	//slice := []int{10, 15, 7, 5, 13, 22, 6, 9}
	fmt.Println(slice)
	//quicksort.QuickSort(slice, quicksort.GetPivot)
	//fmt.Println(slice)
	//fmt.Println(counting_sort.CountingSort(slice))
	//fmt.Println(merge_sort.MergeSort(slice))
	fmt.Println(bucket_sort.BucketSort(slice, bucket_sort.MsBits))

}
