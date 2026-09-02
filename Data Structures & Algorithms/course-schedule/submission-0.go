func canFinish(numCourses int, prerequisites [][]int) bool {
    // adjacency map.
    // key: course id. value: list of courses that *depend on* course id
    neighborMap := make(map[int][]int)
    var buildNeighborMap = func() {
        for _, p := range prerequisites {
            // fmtPrintln(p)
            neighborMap[p[0]] = append(neighborMap[p[0]], p[1])
        }
    }

    buildNeighborMap()
    // fmt.Println(len(neighborMap))
    // 3 states of dfs traversing
    const (
        unvisited = 0
        visiting = 1
        visited = 2
    )
    var checkCycle func(int, []int) bool
    checkCycle = func(node int, state []int) bool {
        //Assigning visiting during recycling
        state[node] = visiting
        // fmt.Fprintln(os.Stdout, node, " neighbormap len ", len(neighborMap[node]))
        for _, nei := range neighborMap[node] {
            if state[nei] == visiting {
                // Found a cycle while in checkCycle recursive dfs run
                return true
            }
            if state[nei] == unvisited {
                // DFS recursive test of dependency graphs of 
                if checkCycle(nei, state) {
                    return true
                }
            }

        }
        state[node] = visited
        return false
    }
    // hasCycle checks of DAG contains atleast 1 cycle
    var hasCycle = func() bool {
        // One of the 3 states
        // TODO: state := make([]int, numCourses)
        // This is actually enough because neighborMap already has the course id as key
        // state := make(map[int]int)
        state := make([]int, numCourses)
        // Check every node
        for k, _ := range neighborMap {
            if state[k] == unvisited {
                if checkCycle(k, state) {
                    return true
                }
            }
        }
        return false
    }

    return !hasCycle()
}