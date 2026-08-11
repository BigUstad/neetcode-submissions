import (
	"slices"
)

type timeEntry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	// A single map pointing directly to a contiguous slice of entries
	store map[string][]timeEntry
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]timeEntry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], timeEntry{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, ok := this.store[key]
	if !ok || len(entries) == 0 {
		return ""
	}

	// BinarySearchFunc allows us to binary search an array of structs
	// We pass a dummy target struct matching our search timestamp
	idx, found := slices.BinarySearchFunc(entries, timeEntry{timestamp: timestamp}, func(a, b timeEntry) int {
		return a.timestamp - b.timestamp
	})

	// Case 1: Exact timestamp match found
	if found {
		return entries[idx].value
	}

	// Case 2: No exact match. BinarySearchFunc returns the index where the 
	// timestamp WOULD be inserted. The largest timestamp less than our target
	// is exactly one index to the left (idx - 1).
	if idx > 0 {
		return entries[idx-1].value
	}

	// Case 3: idx == 0 means the target timestamp is smaller than all existing entries
	return ""
}
