func anagramMappings(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    res := make([]int, len(nums1))
    for i, n := range nums2 {
        m[n] = i
    }
    for i, n := range nums1 {
        res[i] = m[n]
    }
    return res
}
