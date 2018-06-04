package merge_sort

import "fmt"

func doCopy(vals []int) []int {
	tmp := make([]int, len(vals))
	copy(tmp, vals)
	return tmp
}

func MergeSort(items []int) []int {

	if len(items) > 1 {
		mid := len(items) / 2
		left := doCopy(items[0:mid])
		right := doCopy(items[mid:])

		fmt.Println("Left:", left, "Right:", right)

		MergeSort(left)
		MergeSort(right)

		l := 0
		r := 0

		for i := 0; i < len(items); i++ {

			leftValue := -1
			rightValue := -1

			if l < len(left) {
				leftValue = left[l]
			}

			if r < len(right) {
				rightValue = right[r]
			}

			if (leftValue != -1 && rightValue != -1 && leftValue < rightValue) || rightValue == -1 {
				items[i] = leftValue
				l += 1
			} else if (leftValue != -1 && rightValue != -1 && leftValue >= rightValue) || leftValue == -1 {
				items[i] = rightValue
				r += 1
			}

		}
	}
	return items
}
