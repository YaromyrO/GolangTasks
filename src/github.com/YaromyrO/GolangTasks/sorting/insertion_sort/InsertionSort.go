package insertion_sort

func InsertionSort(slice []int) []int {

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
