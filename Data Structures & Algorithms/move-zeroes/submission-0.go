func moveZeroes(nums []int) {
    ln := len(nums)
    i := 0
    for j := 0; j < ln; j++ {
        if nums[j] != 0 {
            nums[i] = nums[j]
            i++
        }
    }
    for i < ln {
        nums[i] = 0
        i++
    }
}
