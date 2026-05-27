package basic

func backtrack(ans *[][]int, num []int, index int) {
	// Base case
	if index >= len(num) {
		temp := make([]int, len(num))
		copy(temp, num)
		*ans = append(*ans, temp)
		return
	}
	for j := index; j < len(num); j++ {
		// swap
		num[index], num[j] = num[j], num[index]
		backtrack(ans, num, index+1)
		// backtrack
		num[index], num[j] = num[j], num[index]
	}
}

func Permute(num []int) [][]int {
	var ans [][]int
	backtrack(&ans, num, 0)
	return ans
}
