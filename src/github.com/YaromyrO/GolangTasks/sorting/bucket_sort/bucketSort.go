package bucket_sort

func MsBits(input []int, item int) int {

	min := input[0]
	max := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		} else if input[i] > max {
			max = input[i]
		}
	}
	return ((item - min) * len(input)) / (max - min + 1)
}

func BucketSort(input []int, getIndex func([]int, int) int) []int {

	if len(input) <= 1 {
		return input
	}

	bucket := make([][]int, len(input))

	for _, v := range input {
		if len(bucket[getIndex(input, v)]) == 0 {
			bucket[getIndex(input, v)] = make([]int, 0)
		}
		bucket[getIndex(input, v)] = append(bucket[getIndex(input, v)], v)
	}

	for _, v := range input {
		if len(bucket[getIndex(input, v)]) < 1 && len(bucket[getIndex(input, v)]) != 0 {
			bucket[getIndex(input, input[v])] = BucketSort(bucket[getIndex(input, input[v])], MsBits)
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
