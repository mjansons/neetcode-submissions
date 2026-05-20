
func getConcatenation(nums []int) (ans []int) {
	ans = append(nums, nums...)
	return ans
}