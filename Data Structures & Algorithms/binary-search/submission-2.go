func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	centre := (left+right)/2

	for left < right {
		if nums[centre] == target {
			return centre
		}
		if nums[centre] > target {
			right = centre - 1
			centre = (left+right)/2
		}else {
			left = centre + 1
			centre = (left+right)/2
		}
	} 

	if nums[centre] == target {
		return centre
	}

	return -1
}
