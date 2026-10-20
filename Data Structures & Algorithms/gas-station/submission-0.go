func canCompleteCircuit(gas []int, cost []int) int {
    gasSum := 0
    costSum := 0
    resIndex := 0

    for _, g := range gas {
        gasSum += g
    }
    total := 0
    for i, c := range cost {
        costSum += c
        total += (gas[i] - cost[i])
        if total < 0 {
            total = 0
            resIndex = i + 1
        }
    }
    if gasSum < costSum {
        resIndex = -1
    }
    if resIndex >= len(cost) {
        resIndex = -1
    }

    return resIndex
}
