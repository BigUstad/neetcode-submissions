import "slices"
func twoSum(numbers []int, target int) []int {
    for i, n := range numbers {
        j, found := slices.BinarySearch(numbers, target-n)
        if found && i != j {
            i, j = min(i, j), max(i, j)
            return []int{i+1,j+1}
        }
    }
    return []int{}
}