package merge_sort

var input []int

func MergeSort(in []int) []int {

	input = in
	doSort(0, len(input)-1)
	return input
}

func doSort(low, high int) {
	if low < high {
		middle := low + (high-low)/2
		doSort(low, middle)
		doSort(middle+1, high)
		mergeSlice(low, middle, high)
	}
}

func mergeSlice(low, middle, high int) {

	i := low
	j := middle + 1
	k := low
	helper := make([]int, len(input))

	copy(helper, input)

	for i <= middle && j <= high {
		if helper[i] <= helper[j] {
			input[k] = helper[i]
			i++
		} else {
			input[k] = helper[j]
			j++
		}
		k++
	}

	for i <= middle {
		input[k] = helper[i]
		k++
		i++
	}
}
