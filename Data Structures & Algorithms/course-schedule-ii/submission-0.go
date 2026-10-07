func findOrder(numCourses int, prerequisites [][]int) []int {
    // I will do it from my understanding of kahn's algorithm
    // Studied when I encountered the task scheduler question in commure interview
    // The adjacency map
    // key : course id. Map of course id to "depends on"
    // adj[k1][k2] = true means k1 must be finished before k2
    adj := make(map[int]map[int]bool)
    // indegree map. Tracks how many prerequisite course, the 'k' course is waiting on
    // indegree[k] = 0 means. You can take course k
    indegree := make(map[int]int)
    // for i := 0; i < numCourses; i++ {
    //     // Yes, redundant assignment
    //     indegree[i] = 0
    //     // Needed
    //     adj[i] = make(map[int]bool)
    // }
    // Build indegree & adjacency maps
    for _, p := range prerequisites {
        indegree[p[0]]++
        if _, exists := adj[p[1]]; !exists {
           adj[p[1]] = make(map[int]bool)
        }
        adj[p[1]][p[0]] = true
    }
    // Build first queue of items with indegree[k] = 0
    // Courses with no prerequisites
    var q []int
    // Resultant courses list, in order that can be completed
    var res []int
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            q = append(q, i)
        }
    }
    if len(q) == 0 { return nil }
    // Count of finished should be numCourses, if all courses can be completed
    finished := 0
    for len(q) > 0 {
        // fmt.Println(q)
        // Front of the queue
        course := q[0]
        q = q[1:]
        finished++
        res = append(res, course)
        for k, _ := range adj[course] {
            indegree[k]--
            if indegree[k] == 0 {
                q = append(q, k)
            }
        }
    }

    if finished != numCourses { return nil }
    return res 
}
