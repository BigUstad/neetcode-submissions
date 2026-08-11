// 1. Fixed constraint to 'any' so it's a true generic Pair
type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

type MatrixIndexPair struct {
	TopLeft     Pair[int, int]
	BottomRight Pair[int, int]
	GridMap     map[int]bool
}

// 2. Declared explicitly as a slice []MatrixIndexPair
var gridLimitPairs = []MatrixIndexPair{
	{
		// 3. Explicitly typed the inner Pair literals
		TopLeft:     Pair[int, int]{First: 0, Second: 0},
		BottomRight: Pair[int, int]{First: 2, Second: 2},
		// 4. Used map literal map[int]bool{} instead of make()
		GridMap:     map[int]bool{}, 
	},
	{
		TopLeft:     Pair[int, int]{First: 0, Second: 3},
		BottomRight: Pair[int, int]{First: 2, Second: 5},
		GridMap:     map[int]bool{},
	},
	{
		TopLeft:     Pair[int, int]{First: 0, Second: 6},
		BottomRight: Pair[int, int]{First: 2, Second: 8},
		GridMap:     map[int]bool{},
	},
	{
		TopLeft:     Pair[int, int]{First: 3, Second: 0},
		BottomRight: Pair[int, int]{First: 5, Second: 2},
		GridMap:     map[int]bool{},
	},
	{
		TopLeft:     Pair[int, int]{First: 3, Second: 3},
		BottomRight: Pair[int, int]{First: 5, Second: 5},
		GridMap:     map[int]bool{},
	},
	{
		TopLeft:     Pair[int, int]{First: 3, Second: 6},
		BottomRight: Pair[int, int]{First: 5, Second: 8},
		GridMap:     map[int]bool{},
	},
	{
		TopLeft:     Pair[int, int]{First: 6, Second: 0},
		BottomRight: Pair[int, int]{First: 8, Second: 2},
		GridMap:     map[int]bool{},
	},
	{
		TopLeft:     Pair[int, int]{First: 6, Second: 3},
		BottomRight: Pair[int, int]{First: 8, Second: 5},
		GridMap:     map[int]bool{},
	},
	{
		TopLeft:     Pair[int, int]{First: 6, Second: 6},
		BottomRight: Pair[int, int]{First: 8, Second: 8},
		GridMap:     map[int]bool{},
	},
}
// 9 row maps
// u2705 Correct and fully allocated
var rowMaps = [9]map[int]bool{
	{}, {}, {}, {}, {}, {}, {}, {}, {},
}
// 9 column maps
var columnMaps = [9]map[int]bool{
	{}, {}, {}, {}, {}, {}, {}, {}, {},
}
func getEleInt(b byte) int {
	ele := int(b)
	if ele == int('.') {
		return -1
	}
	return (ele - '0')
}

func doRowCheck(board [][]byte, i int) bool {
	for x := i; x < 9; x++ {
        // fmt.Print("Doing row check: ")
            // fmt.Print(i)
            // fmt.Print(" , ")
            // fmt.Println(x)
		ele := getEleInt(board[i][x])
		if ele == -1 { continue }
		if rowMaps[i][ele] {
            // fmt.Print("Row check for ")
            // fmt.Print(ele)
            // fmt.Print(" at ")
            // fmt.Print(i)
            // fmt.Print(" , ")
            // fmt.Println(x)
			return false
		}
		rowMaps[i][ele] = true
        columnMaps[x][ele] = true
		if !doGridCheck(board, i, x) {
			return false
		}
	}
	return true
}
func doColumnCheck(board [][]byte, j int) bool {
	for x := j + 1; x < 9; x++ {
        // fmt.Print("Doing column check: ")
            // fmt.Print(x)
            // fmt.Print(" , ")
            // fmt.Println(j)
		ele := getEleInt(board[x][j])
		if ele == -1 { continue }
		if columnMaps[j][ele] {
            // fmt.Print("Col check for ")
            // fmt.Print(ele)
            // fmt.Print(" at ")
            // fmt.Print(x)
            // fmt.Print(" , ")
            // fmt.Println(j)
			return false
		}
        if rowMaps[x][ele] {
            // fmt.Print("Row check in Col check for ")
            // fmt.Print(ele)
            // fmt.Print(" at ")
            // fmt.Print(x)
            // fmt.Print(" , ")
            // fmt.Println(j)
            return false
        }
		rowMaps[x][ele] = true
        columnMaps[j][ele] = true
		// x == j. Skip, added in doRowCheck
		if x != j && !doGridCheck(board, x, j) {
			return false
		}
	}
	return true
}


func doGridCheck(board [][]byte, i int, j int) bool {
	// fmt.Println("Grid check for row: " + strconv.Itoa(i) + " col: " + strconv.Itoa(j))
	for g, l := range gridLimitPairs {
		if i >= l.TopLeft.First && i <= l.BottomRight.First &&
			j <= l.BottomRight.Second && j >= l.TopLeft.Second {
			ele := getEleInt(board[i][j]) // Skipping check
			// Found Grid. Check GridMap.
			if gridLimitPairs[g].GridMap[ele] {
				// fmt.Print("Already Present ")
				// fmt.Print (ele)
				// fmt.Print(" in Grid ")
				// fmt.Println(g)
				return false
			}
			l.GridMap[ele] = true
			break
		}
	}
	return true
}
// i stand in for {i,i}
func doRowColumnCheck(board [][]byte, i int) bool {
	// Add to row map once
	if !doRowCheck(board, i) { return false}
	// Add to column map once
	if !doColumnCheck(board, i) { return false }
	// Add to Grid map once
	return true
}

func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        clear(rowMaps[i])
        clear(columnMaps[i])
        clear(gridLimitPairs[i].GridMap)
    }
	for i := 0 ; i < 9; i++ {
		if !doRowColumnCheck(board, i) {
			return false
		}
	}
	return true
}
