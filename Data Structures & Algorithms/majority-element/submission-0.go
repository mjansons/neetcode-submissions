func majorityElement(nums []int) int {
    numCount := make(map[int]int)
    for _, v := range nums {
        numCount[v]++
    }

	for k, v := range numCount {
		if v > len(nums)/2 {
			return k
		}
	}

	return 0
}
