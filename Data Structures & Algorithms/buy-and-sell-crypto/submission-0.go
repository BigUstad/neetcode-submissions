func maxProfit(prices []int) int {
	bestProfit := 0
	minPrice := math.MaxInt
	for i, p := range prices {
		if p < minPrice { minPrice = p }
		curProfit := (p - minPrice)
		if i > 0 && curProfit > bestProfit {
			bestProfit = curProfit
		}
		// fmt.Println(bestProfit)
	}
	return bestProfit
}
