import "slices"

func findDuplicate(nums []int) int {
    slices.Sort(nums)
    l := len(nums) - 1
    for i, n := range nums {
        if i == l { break }
        if n == nums[i+1] {
            return n
        }
    }
    return 0
}
