/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func getIndexToIndexMaps(head *Node) map[int]int {
	rm := make(map[int]int)
    cm := make(map[*Node]int)
    index := 0
    cur := head
	for cur != nil {
        cm[cur] = index
        cur = cur.Next
        index++
	}
    cur = head
    index = 0
    for cur != nil {
        if randIndex, ok := cm[cur.Random]; ok {
            rm[index] = randIndex
        }
        index++
        cur = cur.Next
    }
    /* fmt.Println()
    if len(cm) != len(rm) {
        fmt.Println("getIndexToIndexMaps lens not same")
    } */
	return rm
}

func getNewLL(head *Node) (map[int]*Node, *Node) {
	cur := head
	index := 0
	newmap := make(map[int]*Node)
	var newHead, newCur, prev *Node
	for cur != nil {
		/* fmt.Print(cur.Val)
		fmt.Print("-> ") */
		newCur = &Node{
			Val : cur.Val,
			Next : nil,
			Random : nil,
		}
		if prev != nil {
			prev.Next = newCur
		}
		newmap[index] = newCur
		if newHead == nil {
			newHead = newCur
			prev = newHead
		}
		cur = cur.Next
		prev = newCur
        index++
	}
    // fmt.Println()
    return newmap, newHead
}

func copyRandomList(head *Node) *Node {
    randmap := getIndexToIndexMaps(head)
	// New map val to cur Node address
    newmap, newHead := getNewLL(head)
    newCur := newHead
    index := 0
    for newCur != nil {
        randIndex, ok := randmap[index]
        if ok {
            newCur.Random = newmap[randIndex]
        }
        newCur = newCur.Next
        index++
    }

	return newHead
}
