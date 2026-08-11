func hasDuplicate(nums []int) bool {
    if len(nums) == 0 {
        return false
    }
    duplMap := make(map[int]bool)
    for _, n:= range nums {
        if duplMap[n] {
            return true
        }
        duplMap[n] = true
    }
    return false
}
