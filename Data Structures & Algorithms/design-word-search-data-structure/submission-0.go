type TrieNode struct {
    children [26]*TrieNode
    eow bool
    // has no meaning for root. Only the children
    length uint8
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    var w WordDictionary
    w.root = &TrieNode{}
    w.root.eow = false
    return w
}

func (this *WordDictionary) AddWord(word string)  {
	parent := this.root
	l := len(word)
	for i, c := range word {
		idx := c - 'a'
		if parent.children[idx] == nil {
			parent.children[idx] = &TrieNode{}
		}
		if i == l-1 {
			parent.children[idx].eow = true
		}
		parent = parent.children[idx]
        // fmt.Print(word, " cur ")
        // fmt.Fprintf(os.Stdout, "%c ", word[i])
        // fmt.Println(parent.length)
	}
}

func (this *WordDictionary) Search(word string) bool {
    l := len(word)
    found := false
    var dfs func(int, *TrieNode) bool
    // dfs is the correct answer
    dfs = func(start int, parent *TrieNode) bool {
        // fmt.Println("dfs ", word[start:])
        for i := start; i < l; i++ {
            c := word[i]
            if c == '.' {
                for _, child := range parent.children {
                    if child != nil && dfs(i+1, child) {
                        return true
                    }
                }
                // fmt.Println("1. Nope")
                return false
            } else {
                // Other normal processing
                idx := c - 'a'
                if parent.children[idx] == nil {
                    // fmt.Println("2. Nope")
                    return false
                }
                parent = parent.children[idx]
            }
        }
        return parent.eow
    }
    found = dfs(0, this.root)
    return found
}
