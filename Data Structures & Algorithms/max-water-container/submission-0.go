func maxArea(heights []int) int {
    start, end := 0, (len(heights) - 1)
    prevAreaMax, curArea, curAreaMax := 0, 0, 0

    // Lets move the indices after area calculations
    // Two pointers approach. Earlier the better as bottom needs to be maximized
    for start < end {
        curArea = min(heights[start], heights[end]) * (end - start)
        prevAreaMax = curAreaMax
        curAreaMax = max(curArea, prevAreaMax)
        if heights[start] < heights[end] {
            start++
        } else {
            end--
        }
    }
    return curAreaMax
}
