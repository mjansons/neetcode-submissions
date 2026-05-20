func twoSum(nums []int, target int) []int {
	var myMap map[int]int = make(map[int]int)

	for i, v := range nums {
		diff := target - v
		_, exists := myMap[diff]
		if exists {
			return []int{ myMap[diff], i}
		}
		myMap[v] = i
	}
	return []int{}
}