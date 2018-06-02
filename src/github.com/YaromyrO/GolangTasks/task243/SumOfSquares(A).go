package task243

func GetPair(number int) []int {
	var result []int
	for i := 1; i*i <= number; i++ {
		for j := 1; j*j <= number; j++ {
			if i*i+j*j == number {
				result = append(result, i, j)
				return result
			}
		}
	}
	return make([]int, 0)
}
