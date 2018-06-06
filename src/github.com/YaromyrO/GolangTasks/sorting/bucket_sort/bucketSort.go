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
		index := getIndex(input, v)
		bucket[index] = append(bucket[index], v)
		sort.Ints(bucket[index])
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
