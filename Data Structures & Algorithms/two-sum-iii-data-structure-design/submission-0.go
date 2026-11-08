import("slices")
// The leetcode solution I've is straightforward
// I'd like to implemented the sorted solution in Neetcode
type TwoSum struct {
	keys []int
	sorted bool
}
func Constructor() TwoSum{
	return TwoSum{
		sorted: false,
	}
}

// When adding a key, add it to the end
// Given that pdqsort in go Slices will ensure that subsequent sorts are quicker
func(this *TwoSum)Add(number int){
	this.keys = append(this.keys, number)
	this.sorted = false
}
// Lazy initialization. Sort is done only when find is needed
func(this *TwoSum)Find(value int)bool{
	if !this.sorted {
		slices.Sort(this.keys)
		this.sorted = true
	}
	for _, k1 := range this.keys {
		k2 := value - k1
		fi, found := slices.BinarySearch(this.keys, k2)
		if !found { continue }
		// k2 found
		if k1 != k2 { return true }
		// If we have 2 instances of k1 & k2, that satisfies the condition
		if fi < (len(this.keys) - 1) && this.keys[fi+1] == k2 {
			return true
		}
	}
	return false
}
