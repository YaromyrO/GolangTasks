package quicksort

func GetPivot(input []int) int {

	if len(input) == 0{
		return 0
	}

	return input[(len(input)-1) / 2]
}

func QuickSort(input []int, p func([]int) int) {

	if len(input) >= 2 {

		left := make([]int, 0)
		right := make([]int, 0)
		middle := make([]int, 0)
		pivot := p(input)

		for _, v := range input {
			if v < pivot {
				left = append(left, v)
			}
			if v == pivot {
				middle = append(middle, v)
			}
			if v > pivot {
				right = append(right, v)
			}
		}

		QuickSort(left, GetPivot)
		QuickSort(right, GetPivot)

		left = append(left, middle...)
		left = append(left, right...)

		copy(input, left)
	}
}

