func findJudge(n int, trust [][]int) int {
    // You could Calculate indegree counts
    indegree := make([]int, n+1)
    rIndegree := make([]int, n+1)
    for _, t := range trust {
        indegree[t[1]]++
        rIndegree[t[0]]++
    }
    for i := 1; i <= n; i++ {
        if indegree[i] == (n-1) &&
            rIndegree[i] == 0 {
            return i
        }
    }
    return -1
}
