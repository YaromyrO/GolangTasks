package quicksort

func getPivot(input []int) int {

	return input[(len(input) / 2)]
}

func QuickSort(input []int, f func([]int) int) []int {

	if len(input) <= 1 {
		return input
	}

	low := 0
	high := len(input) - 1
	i := low
	j := high
	pivot := getPivot(input)

	if i < j {
		for input[i] < pivot {
			i++
		}
		for input[j] > pivot {
			j--
		}
		if i <= j {
			input[i], input[j] = input[j], input[i]
			i++
			j--
		}
	}

	if low < j {
		QuickSort(input[low:j+1], nil)
	}
	if i < high {
		QuickSort(input[i:high+1], nil)
	}

	return input
}
