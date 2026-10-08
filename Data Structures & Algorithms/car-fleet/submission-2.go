import (
	"slices"
	"cmp"
)

// No stack solution here.
// Process cars in descending order of position (closest to the target first).
// Remember the arrival time of the last fleet.
// O(n) time. O(1) space
func carFleet(target int, position []int, speed []int) int {
    n := len(position)

    pairs := make([][2]int, n)
    for i := 0; i < n; i++ {
        pairs[i] = [2]int{position[i], speed[i]}
    }

    slices.SortFunc(pairs, func(a, b [2]int) int {
        return cmp.Compare(b[0], a[0])
    })

    fleetCount := 0
    var prevTime float64

    for _, p := range pairs {
        tt := float64(target-p[0]) / float64(p[1])

        if fleetCount == 0 || tt > prevTime {
            fleetCount++
            prevTime = tt
        }
    }

    return fleetCount
}

// Stack solution here
/*func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	// Pair each car's position with its speed
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{position[i], speed[i]}
	}
	// Sort the cars in descending order of position,
	// closest to target appearing first
	slices.SortFunc(pairs, func(a, b [2]int) int{
		return cmp.Compare(b[0], a[0])
	})
	fleetStack := make([]float64, 0, n)
	for _, p := range pairs {
		// Compute time taken to reach target
		tt := float64(target - p[0]) / float64(p[1])
		// If the length is more than 1, find out
		// if the new car's time taken is less than or equal
		// to time-taken to the one before it
		// That means it catches up & merges with that fleet.
		// Hence pop it with the stack
		// ls - last of stack
		if len(fleetStack) == 0 || tt > fleetStack[len(fleetStack)-1] {
            fleetStack = append(fleetStack, tt)
        }
		// fmt.Println(fleetStack, ", ", len(fleetStack))
        // This is a monotonic stack solution. The stack stores fleet arrival times, not car positions.
        // Cars are processed from closest to farthest.
        // tt pushed to stack represents a fleet that the next car cannot catch up with.
	}
	// Remaining times in the stack eventually give the count of fleet
	return len(fleetStack)
}*/
