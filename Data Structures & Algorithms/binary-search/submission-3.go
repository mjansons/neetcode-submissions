func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	
	for left <= right {
		centre := (left+right)/2
		if nums[centre] == target {
			return centre
		}
		if nums[centre] > target {
			right = centre - 1
		}else {
			left = centre + 1
		}
	} 
	return -1
}
