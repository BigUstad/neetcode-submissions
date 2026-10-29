import("slices")
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    // NOTE: This is an implementation of Bellman-Ford
    // Read up on Bellman-Ford, Dijktra's solutions.
    // Also read up on the most optimized Shortest Path Faster algorithm
    // 1. Initialize prices array with infinity, set prices[src] = 0.
    prices := make([]int, n)
    for i := 0; i < n; i++ {
        prices[i] = math.MaxInt
    }
    prices[src] = 0
    // 2. Repeat k + 1 times:
    for i := 0; i <= k; i++ {
        //     • Make a copy tmpPrices = prices.
        tmpPrices := slices.Clone(prices)
        //     • For every flight (s -> d, cost):
        for _, f := range flights {
            // source dest & price
            s, d, p := f[0], f[1], f[2]
            if prices[s] == math.MaxInt {
                continue
            }
            //         • If prices[s] is reachable:
            //         • Try relaxing the edge: tmpPrices[d] = min(tmpPrices[d], prices[s] + cost)
            tmpPrices[d] = min(tmpPrices[d], prices[s] + p)
        }
        // Just releasing the memory before assignment
        //     • Assign prices = tmpPrices.
        prices = tmpPrices
        // No minheap here in this problem. - So is there another problem where minheap is needed?
    }
    // 3. After all iterations:
    //     • If prices[dst] is still infinity, return -1.
    //     • Otherwise, return prices[dst].
    if prices[dst] == math.MaxInt {
        return -1
    }
    return prices[dst]
}
