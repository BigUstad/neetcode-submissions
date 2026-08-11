func search(nums []int, target int) int {
    if len(nums) == 1 && nums[0] == target {
        return 0
    }
    if len(nums) == 1 && nums[0] != target {
        return -1
    }
    if len(nums) == 2 {
        if nums[0] == target {
            return 0
        } else if nums[1] == target {
            return 1
        } else {
            return -1
        }
    }
    left := 0
    right := len(nums) - 1
    for left <= right {
        m := left + ((right - left) / 2)
        if nums[m] < target {
            left = m + 1
        } else if nums[m] > target {
            right = m - 1
        } else {
            return m
        }
    }
    return -1
}
