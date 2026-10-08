import (
	"slices"
	"cmp"
)

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	// Pair each car's position with its speed
	var pairs [][2]int
	for i := 0; i < n; i++ {
		pairs = append(pairs, [2]int{position[i], speed[i]})
	}
	// Sort the cars in descending order of position,
	// closest to target appearing first
	slices.SortFunc(pairs, func(a, b [2]int) int{
		return cmp.Compare(b[0], a[0])
	})
	var fleetStack []float64
	for _, p := range pairs {
		// Compute time taken to reach target
		tt := float64(target - p[0]) / float64(p[1])
		fleetStack = append(fleetStack, tt)
		// If the length is more than 1, find out
		// if the new car's time taken is less than or equal
		// to time-taken to the one before it
		// That means it catches up & merges with that fleet.
		// Hence pop it with the stack
		// ls - last of stack
		ls := len(fleetStack) - 1
		if ls >= 1 && fleetStack[ls] <= fleetStack[ls-1] {
			fleetStack = fleetStack[:ls]
		}
		// fmt.Println(fleetStack, ", ", len(fleetStack))
	}
	// Remaining times in the stack eventually give the count of fleet
	return len(fleetStack)
}
