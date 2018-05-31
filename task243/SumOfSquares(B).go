package main

func GetPairs(number int) []int {
	var result []int
	for i := 1; i * i <= number; i++ {
		for j := 1; j * j <= number; j++ {
			if i * i + j * j == number && i >= j{
				result = append(result, i, j)
			}
		}
	}
	if result != nil {
		return result
	}
	return make([]int, 0)
}