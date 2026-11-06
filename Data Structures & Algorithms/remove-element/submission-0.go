func removeElement(nums []int, val int) int {
    ln := len(nums)
    i := 0
    for j := 0; j < ln; j++  {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}
