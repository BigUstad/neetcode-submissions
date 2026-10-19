func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // This is an implementation of the most optimal Binary search solution.
    // The Two pointer solution that keeps track of median1 & median2
    // Seemed more intuitive to me. But it is O(n+m) time.
    n1 := len(nums1)
    n2 := len(nums2)
    total := n1 + n2
    isOddLength := (total % 2 != 0)
    // Why did neetcode make the adjustment of total+1
    var half int = ((total+1) / 2)
    var median float64 = 0.0
    if n1 > n2 {
        nums1, nums2 = nums2, nums1
        n1, n2 = n2, n1
    }
    // fmt.Println(nums1, n1)
    // Now nums1 is ensured to be always smaller.
    left, right := 0, n1
    // Time - O(Log min(n, m))
    for left <= right {
        // m + n guaranteed to be 2000, so skipping the overflow check based equation
        // Binary search on nums1. i is the cut index in nums1, mid for left & right
        var i int = (left + right)/2 // # nums1 pointer
        // fmt.Println(left, right, i)
        j := half - i // index of the mid-point.
        // Determine if we have found the median
        // We define border values around the cut index (i)
        n1Left := math.MinInt64
        if i > 0 {
            n1Left = nums1[i-1]
        }
        n1Right := math.MaxInt64
        if i < n1 {
            n1Right = nums1[i]
        }
        n2Left := math.MinInt64
        if j > 0 {
            n2Left = nums2[j-1]
        }
        n2Right := math.MaxInt64
        if j < n2 {
            n2Right = nums2[j]
        }
        // Left partition is quite correct
        // n1 pointer has to keep approaching end of nums1 array
        // n2 pointer has to keep approaching the beginning of nums2 array
        // fmt.Println(n1Left, n1Right, n2Left, n2Right)
        if n1Left <= n2Right && n2Left <= n1Right {
            // Odd or Even length, determines different medians
            if isOddLength {
                median = float64(max(n1Left, n2Left))
            } else {
                sum := max(n1Left, n2Left) + min(n1Right, n2Right)
                median = float64(sum) / float64(2.0)
            }
            return median
        } else if n1Left > n2Right {
            right = i - 1
        } else {
            left = i + 1
        }
    }
    return median
}