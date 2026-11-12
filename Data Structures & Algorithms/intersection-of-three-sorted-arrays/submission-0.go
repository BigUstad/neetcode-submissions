import("slices")
func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    l1, l2, l3 := len(arr1), len(arr2), len(arr3)
	// Pick the smallest array
	minLn := min(l1, l2, l3)
	var res []int
	other1 := arr2
	other2 := arr3
	smallest := arr1
	if l2 == minLn {
		smallest = arr2
		other1 = arr1
		other2 = arr3
	} else if l3 == minLn {
		smallest = arr3
		other1 = arr1
		other2 = arr2
	}
	for _, s := range smallest {
		_, f1 := slices.BinarySearch(other1, s)
		_, f2 := slices.BinarySearch(other2, s)
		if f1 && f2 {
			res = append(res, s)
		}
	}
	return res
}
