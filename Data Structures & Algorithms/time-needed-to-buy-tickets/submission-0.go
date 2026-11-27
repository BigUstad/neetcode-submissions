func timeRequiredToBuy(tickets []int, k int) int {
    ln := len(tickets)
    q := make([]int, ln)
    for i, _ := range tickets {
        q[i] = i
    }
    i := 0
    front := 0
    for len(q) > 0 && tickets[k] > 0 {
        tickets[q[front]]--
        fe := q[0]
        if tickets[fe] > 0 {
            q = append(q, fe)
        }
        // fmt.Println(q)
        q = q[1:]
        i++
    }
    return i
}