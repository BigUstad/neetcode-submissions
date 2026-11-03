func search(nums []int, target int) bool {
	n := len(nums)
	low := 0
	high := n - 1
	if n == 0 { return false }
	if n == 1 {
		if nums[0] == target { return true }
		return false
	}
	for low <= high {
		mid := low + (high - low)/2
		if nums[mid] == target {
			return true
		}
		// Step 1. This would mean, left part is correctly sorted & rotation has ended
		if nums[low] < nums[mid] {
			// Check if the target is in sorted left half
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1 // Go left
			} else {
				// Check for the target in the sorted right half
				low = mid + 1 // Go right
			}
		} else if nums[low] > nums[mid] {
			// Step 2. This would mean the left part is not sorted.
			// Right half is sorted somewhat. Atleast mid to high
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1 // Go right
			} else {
				high = mid - 1 // Go left
			}
		} else {
			// nums[low] == nums[mid]
			// Let's skip the duplicate & search the part where it is unique numbers
			low++
		}
	}
	return false
}

