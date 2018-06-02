package counting_sort

func CountingSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	min := slice[0]
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		} else if slice[i] > max {
			max = slice[i]
		}
	}

	result := make([]int, len(slice))
	counts := make([]int, max - min+1)

	for _,v := range slice {
		counts[v-min]++
	}

	z := 0
	for i, c := range counts {
		for ; c > 0; c-- {
			result[z] = i + min
			z++
		}
	}
	return result
}
