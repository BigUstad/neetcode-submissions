func subsetXORSum(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }

    totalSum := 0

    var backtrack func(start int, currentXOR int)
    backtrack = func(start int, currentXOR int) {
        // Add the XOR total of the current subset to our global total
        totalSum += currentXOR

        // Explore further subsets by including remaining elements
        for i := start; i < n; i++ {
            // Calculate the next running XOR value on the fly
            backtrack(i+1, currentXOR ^ nums[i])
        }
    }

    // Start backtracking with the first index and an initial XOR sum of 0
    backtrack(0, 0)

    return totalSum
}
