type TrieNode struct {
	children [26]*TrieNode
	// eow - end of word
	eow bool
    // word at the node
    word string
	// root - true for only one node
	isRoot bool
}

type Trie struct {
	root *TrieNode
}

func constructInputTrie(words []string) *Trie {
    var res Trie
    res.root = &TrieNode{
        eow: false,
        isRoot: true,
    }
    for _, w := range words {
        cur := res.root
        for _, c := range w {
            // index within cur trie node where the child char goes
            i := c - 'a'
            if cur.children[i] == nil {
                cur.children[i] = &TrieNode{}
            }
            cur = cur.children[i]
        }
        cur.eow = true
        cur.word = w
    }
    return &res
}

func findWords(board [][]byte, words []string) []string {
    var res []string
	if len(board) == 0 { return res }
	if len(words) == 0 { return res }

    inTrie := constructInputTrie(words)
    rows, cols := len(board), len(board[0])
    visit := make([][]bool, rows)
    for r := range visit {
        visit[r] = make([]bool, cols)
    }
    var dfs func(int, int, *TrieNode)

    dfs = func(r, c int, node *TrieNode) {
        // Ending base case
        //   row or column out of bounds
        //   Already visited
        //   *this* node doesn't have the children nodes for the character set.
        //   NOTE: node is based on trie built out of input words
        if r < 0 || c < 0 ||
            r == rows || c == cols ||
            visit[r][c] || node.children[board[r][c] - 'a'] == nil {
            // fmt.Println("2.", node.word)
            return
        }
        // Marking as visited
        visit[r][c] = true
        node = node.children[board[r][c] - 'a']
        // Important end case
        // Check uniqueness, so no word gets added twice
        if node != nil && node.eow && len(node.word) > 0 {
            // fmt.Println("1.", node.word)
            res = append(res, node.word)
            // nullifying node.word because it has already been added
            node.word = ""
        }
        dfs(r - 1, c, node)
        dfs(r + 1, c, node)
        dfs(r, c - 1, node)
        dfs(r, c + 1, node)
        // Backtracking move - if done processing, remove visited entry
        visit[r][c] = false
    }
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            idx := board[r][c] - 'a'
            if inTrie.root.children[idx] != nil {
                dfs(r, c, inTrie.root)
            }
        }
    }
    return res
}
