func firstMissingPositive(nums []int) int {
	// clean
	for i ,v := range nums {
		if v < 1 || v > len(nums) {
			nums[i] = len(nums) + 1
		}
	}
	// check presence with negation
	for i := range nums{
		tmp := nums[i]
		if tmp < 0 {
			tmp = -tmp
		}
		if tmp == len(nums) + 1 {
			continue
		}

		if nums[tmp - 1] > 0 {
		nums[tmp - 1] = -nums[tmp - 1]
		}

	}
	// find first
	for i := range len(nums) {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
