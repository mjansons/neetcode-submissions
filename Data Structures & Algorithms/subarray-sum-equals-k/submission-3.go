func subarraySum(nums []int, k int) int {

	result := 0
	seenSum := map[int]int{
		0 : 1,
	}
	currentSum := 0

	for _, v := range nums {
		currentSum += v

		result += seenSum[currentSum - k]

		seenSum[currentSum]++
	}
	return result
}
