func subarraySum(nums []int, k int) int {

	sum := 0
	for i, v := range nums {
		internalSum := v

		if internalSum == k {
			sum++
		}
		for _, v2 := range nums[i+1:] {
			internalSum += v2
			if internalSum == k {
				sum++
			}
		}
	}
	return sum
}
