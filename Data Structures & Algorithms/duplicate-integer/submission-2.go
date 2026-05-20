func hasDuplicate(nums []int) bool {
    var seen map[int]int = make(map[int]int)

    for _ , v := range nums {
        _, exists := seen[v]
        if !exists {
            seen[v] = 1
            continue
        }
        return true
    }

    return false
}
