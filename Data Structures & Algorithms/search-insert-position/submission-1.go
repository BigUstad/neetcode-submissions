import ("slices")
func searchInsert(nums []int, target int) int {
    j, _ := slices.BinarySearch(nums, target)
    return j
}
