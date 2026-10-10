func minCostClimbingStairs(cost []int) int {
    l := len(cost) - 3
    // Changing the array in place
    // Calculating the minimum for the step i
    // Last index is l - 1.
    // Starting with l - 3 as you can get to l - 2 & l - 1.
    for i := l ; i >= 0; i-- {
        cost[i] += min(cost[i+1], cost[i+2])
    }
    return min(cost[0], cost[1])
}
