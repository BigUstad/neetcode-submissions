import (
  pq "github.com/emirpasic/gods/queues/priorityqueue"
  "github.com/emirpasic/gods/utils"
)

func maxIntComparator(a, b interface{}) int {
	return -utils.IntComparator(a.(int), b.(int))
}

func lastStoneWeight(stones []int) int {
	// edge cases
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	if len(stones) == 2 {
		return int(math.Abs(float64(stones[1] - stones[0])))
	}
	queue := pq.NewWith(maxIntComparator)
	for _, s := range stones {
		queue.Enqueue(s)
	}
	for !queue.Empty() {
		s1Ele, _ := queue.Dequeue()
		s1 := s1Ele.(int)
		s2Ele, ok := queue.Dequeue()
		if !ok { return s1 }
		s2 := s2Ele.(int)
		if s1 != s2 {
			queue.Enqueue(int(math.Abs(float64(s2 - s1))))
		}
	}
	return 0
}
