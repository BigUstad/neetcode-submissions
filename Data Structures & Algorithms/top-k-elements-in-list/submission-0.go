func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	n := len(nums)
	for _, i := range nums {
		numMap[i]++
	}
	numBuckets := make([][]int,n+1)
	// Bucket indexed
	for k, v := range numMap {
		numBuckets[v] = append(numBuckets[v], k)
	}
	var topk []int
	tillk := 0
	for i := n; i > 0; i-- {
		if len(numBuckets[i]) == 0 {
			continue
		}
		tillk += len(numBuckets[i])
		if tillk > k {
			break
		}
		topk = append(topk, numBuckets[i]...)
	}
	// fmt.Println(len(topk))
	// fmt.Println(tillk)
	return topk[:k]
}
