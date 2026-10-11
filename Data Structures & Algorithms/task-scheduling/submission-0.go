import (
  pq "github.com/emirpasic/gods/queues/priorityqueue"
  "github.com/emirpasic/gods/utils"
)

type Element struct {
    count int
    time int
}

func leastInterval(tasks []byte, n int) int {
    // The taskmap keeps a count of each task in the input tasks queue
    taskMap := make(map[byte]int)
    for _, t := range tasks {
        taskMap[t]++
    }
    // Max heap keeps each individual task's metadata
    // And count of each task
    maxQ := pq.NewWith(func(i, j interface{}) int {
        return -utils.IntComparator(i, j)
    })
    for _, v := range taskMap {
        maxQ.Enqueue(v)
    }
    // Empty FIFO queue. Stores remaining count after run,
    // and next time of run
    q := make([]Element, 0)
    // time initialized to 0 before starting processing
    time := 0
    // Keep executing while max heap is not empty
    // And elements in cooldown queue are not all dealt with
    for !maxQ.Empty() || len(q) > 0 {
        // fmt.Println("max-heap ", maxQ)
        // fmt.Println("cooling-queue", q)
        // Incrememnt queue
        time++
        // I removed the maxQ.Empty() if condition
        // I'd prefer to increment time till front of the queue
        // time is accrued
        if !maxQ.Empty() {
            // If heap is not empty. Pop the task with largest remaining count
            // Decrement count. Marking it as executed once.
            // Only if count is greater than 0, push it to cooling queue,
            // with time updated for future run
            top, _ := maxQ.Dequeue()
            cnt := top.(int) - 1
            if cnt > 0 {
                q = append(q, Element{count: cnt, time: time + n})
            }
        }
        // Check front of cooldown queue
        // "Execute" front of queue if current time is same as time of front
        // Enqueue in max-heap the remaining count
        // Pop the front of cooling queue
        if len(q) > 0 && time == q[0].time {
            maxQ.Enqueue(q[0].count)
            q = q[1:]
        }
    }

    // When both heap & cooling queue are empty, 'time' has the minimum time required to finish all tasks
    return time
}
