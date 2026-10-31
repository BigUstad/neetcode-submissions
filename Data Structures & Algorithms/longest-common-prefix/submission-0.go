type TrieNode struct {
    // Set the char
    children [26]*TrieNode
    // end of word
    eow bool
    // root - true for only the top node.
    isRoot bool
    // how many times is this node set
    setCount int
}

type Trie struct {
    root *TrieNode
}

func Constructor() *Trie {
    return &Trie{
        root: &TrieNode{isRoot: true},
    }
}

func (t *Trie) Insert(word string) {
    parent := t.root
    for _, c := range word {
        idx := c - 'a'
        if parent.children[idx] == nil {
            parent.children[idx] = &TrieNode{}
        }
        parent.children[idx].setCount++
        parent = parent.children[idx]
    }
    // parent is now assigned to the last alphabet
    parent.eow = true
}

func longestCommonPrefix(strs []string) string {
    var prefix string
    sl := len(strs)
    trie := Constructor()
    for i := 0; i < sl; i++ {
        trie.Insert(strs[i])
    }
    startPrefix := -1
    for i, c := range trie.root.children {
        if c != nil && c.setCount == sl {
            // Found first startCount
            startPrefix = i
            break
        }
    }
    // fmt.Println("s = ", startPrefix)
    if startPrefix == -1 { return "" }
    prefix = string('a'+byte(startPrefix))
    cur := trie.root.children[startPrefix]
    for cur != nil && !cur.eow {
        prevCur := cur
        for i, c := range cur.children {
            if c != nil && c.setCount == sl {
                cur = c
                prefix += string('a'+byte(i))
                break
            }
        }
        // fmt.Println("cur = ", prefix)
        if prevCur == cur { break }
    }
    return prefix
}
