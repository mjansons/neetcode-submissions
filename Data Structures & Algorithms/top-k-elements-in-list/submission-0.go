func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	bucket := make ([][]int, len(nums) + 1)
	var result []int

	for _, v := range nums {
		count[v]++
	}

	for key, value := range count {
		bucket[value] = append(bucket[value], key)
	}

	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 && len(result) < k {
			for _, v := range bucket[i] {
				result = append(result, v)

			}
		}
	}
	return result
}
