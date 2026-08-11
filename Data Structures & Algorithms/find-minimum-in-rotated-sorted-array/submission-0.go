func findMin(nums []int) int {
    n := len(nums)
    if n == 0 { return -1 }
    if n == 1 { return nums[0] }
    if n == 2 {
        if nums[0] < nums[1] {
            return nums[0]
        }
        return nums[1]
    }
    low := 0
    high := n - 1 
    for low <= high {
        mid := low + (high - low) / 2
        if mid == 0 || mid == (n - 1) { break }
        // Inflection point where mid is the index of lowest ele
        // fmt.Println(mid)
        if nums[mid] < nums[mid + 1] &&
            nums[mid - 1] > nums[mid] {
            return nums[mid]
        }
        if nums[mid] > nums[high] {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return min(nums[low], nums[high])
}
