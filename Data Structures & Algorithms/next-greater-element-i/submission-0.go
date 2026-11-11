func nextGreaterElement(nums1 []int, nums2 []int) []int {
    n2m := make(map[int]int)
    for i, n := range nums2 {
        n2m[n] = i
    }
    var stck []int
    var res []int
    // l1 := len(nums1)
    l2 := len(nums2)
    ms2 := make([]int, l2)
    for i := 0; i < l2; i++ {
        ms2[i] = -1
    }
    for i := 0; i < l2; i++ {
        var top int
        ls := len(stck)
        if ls > 0 {
            top = stck[ls-1]
        }
        for ls > 0 && nums2[top] < nums2[i] {
            ms2[top] = nums2[i]
            stck = stck[:ls-1]
            ls--
            if ls > 0 {
                top = stck[ls-1]
            }
        }
        stck = append(stck, i)
    }
    // fmt.Println(ms2)
    for _, n := range nums1 {
        res = append(res, ms2[n2m[n]])
    }

    return res
}
