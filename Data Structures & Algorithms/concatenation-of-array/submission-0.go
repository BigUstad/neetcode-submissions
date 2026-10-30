func getConcatenation(nums []int) []int {
    l := len(nums)
    newNums := make([]int, 2 * l)
    for i, n := range nums {
        newNums[i] = n
        newNums[i+l] = n
    }
    return newNums
}
