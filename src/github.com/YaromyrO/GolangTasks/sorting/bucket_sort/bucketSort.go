package bucket_sort

import "sort"

func MsBits(input []int, item int) int {

	max := input[0]
	min := input[0]
	for _, value := range input {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return ((item - min) * len(input)) / (max - min + 1)
}

func BucketSort(input []int, getIndex func([]int, int) int) []int {

	if len(input) <= 1 {
		return input
	}

	bucket := make(map[int][]int)

	for _, v := range input {
		if len(bucket[getIndex(input, v)]) == 0 {
			bucket[getIndex(input, v)] = make([]int, 0)
		}
		bucket[getIndex(input, v)] = append(bucket[getIndex(input, v)], v)
	}

	for _, v := range input {
		if len(bucket[getIndex(input, v)]) > 1 {
			//bucket[getIndex(input, v)] = insertionSort(bucket[getIndex(input, v)])
			sort.Ints(bucket[getIndex(input, v)])
		}
	}



	x := 0
	for i := 0; i < len(input); i++ {
		for _, v := range bucket[i] {
			input[x] = v
			x++
		}
	}
	return input
}

func insertionSort(slice []int) []int {

	for i := 1; i < len(slice); i++ {
		temp := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > temp {
			slice[j+1] = slice[j]
			j = j - 1
		}
		slice[j+1] = temp
	}
	return slice
}
