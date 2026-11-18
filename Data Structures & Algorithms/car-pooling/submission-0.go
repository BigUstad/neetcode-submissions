import(
	"slices"
	"cmp"
)

func carPooling(trips [][]int, capacity int) bool {
    if len(trips) == 0 {
        return true
    }
    if len(trips) == 1 {
        return trips[0][0] <= capacity
    }
    slices.SortFunc(trips, func(a, b []int) int {
        return cmp.Or(
            cmp.Compare(a[2], b[2]),
            cmp.Compare(a[1], b[1]),
        )
    })
    // So that's sorted out :-)
    // Get the heighest 'to'
    highTo := trips[len(trips)-1][2]
    diffArray := make([]int, highTo+2)
    for _, t := range trips {
        // Passengers board the car
        diffArray[t[1]] += t[0]
        // Passengers will disembark at t[2]
        diffArray[t[2]] -= t[0]
    }
    // Sweep through the array
    prev := 0
    for i := 0; i <= highTo; i++ {
        diffArray[i] += prev
        prev = diffArray[i]
    }
    // fmt.Println(capacity, diffArray)
    return slices.Max(diffArray) <= capacity
}
