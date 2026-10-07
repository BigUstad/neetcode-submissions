type TrieNode struct {
	children [26]*TrieNode
	// eow - end of word
	eow bool
	// root - true for only one node
	isRoot bool
}
type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{root: &TrieNode{isRoot: true}}
}

func (this *PrefixTree) Insert(word string) {
	parent := this.root
	l := len(word) - 1
	for i, c := range word {
		idx := c - 'a'
		if parent.children[idx] == nil {
			parent.children[idx] = &TrieNode{}
		}
		if i == l {
			// fmt.Println(word, "last", c)
			parent.children[idx].eow = true
		}
		parent = parent.children[idx]
	}
}

func (this *PrefixTree) Search(word string) bool {
	parent := this.root
	l := len(word) - 1
	foundWord := false
	for i, c := range word {
		idx := c - 'a'
		if parent.children[idx] == nil {
			return false
		}
		if i == l && parent.children[idx].eow {
			foundWord = true
		}
		parent = parent.children[idx]
	}
	return foundWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	parent := this.root
	l := len(prefix) - 1
	foundWord := false
	for i, c := range prefix {
		idx := c - 'a'
		if parent.children[idx] == nil {
			return false
		}
		if i == l {
			foundWord = true
		}
		parent = parent.children[idx]
	}
	return foundWord	
}
