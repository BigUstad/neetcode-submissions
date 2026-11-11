func largestUniqueNumber(nums []int) int {
    m := make([]int, 1001)
    for _, n := range nums {
        m[n]++
    }
    res := -1
    for _, n := range nums {
        if m[n] == 1 && n > res {
            res = n
        }
    }
    return res
}
