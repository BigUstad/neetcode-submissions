func majorityElement(nums []int) []int {
	//TODO: Boyer-Moore Hashmap Voting algorithm.
	// Bit complicated but probably more useful in generic cases
    // Also, Frequency count array is a good O(n) space solution
    // This is a Boyer Moore voting algorithm implementation with 2 candidates
    // Assuming 2 candidates would take up more than 2/3 of the array space
    num1, num2 := math.MinInt, math.MinInt
    count1, count2 := 0, 0
    expectedCount := len(nums) / 3

    for _, n := range nums {
        if count1 == 0 && n != num2 {
            num1 = n
            count1 = 1
            continue
        } else if count2 == 0 && n != num1 {
            num2 = n
            count2 = 1
            continue
        }
        if n == num1 {
            count1++
        } else if n == num2 {
            count2++
        } else {
            if count1 > 0 { count1-- }  
            if count2 > 0 { count2-- }
        }
        // fmt.Println(num1, count1)
        // fmt.Println(num2, count2)
        // fmt.Println()
    }
    // fmt.Println(num1, count1)
    // fmt.Println(num2, count2)
    // fmt.Println()
    count1, count2 = 0, 0
    for _, n := range nums {
        if n == num1 {
            count1++
        }
        if n == num2 {
            count2++
        }
    }
    var res []int
    if count1 > expectedCount {
        res = append(res, num1)
    }
    if count2 > expectedCount {
        res = append(res, num2)
    }
    return res
}
