func removeElement(nums []int, val int) int {
	lastIndex := 0
	seenVal := false
	k := 0

	if len(nums) > 0 && nums[0] == val {
		lastIndex = -1
	}

	for i := range nums {
		if nums[i] != val {
			if seenVal {
				lastIndex++
				nums[lastIndex] = nums[i]
				continue
			}
			lastIndex = i
			continue
		}
		seenVal = true
		k++
	}
	return len(nums)-k
}
