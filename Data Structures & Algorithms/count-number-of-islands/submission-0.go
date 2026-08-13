import llq "github.com/emirpasic/gods/queues/linkedlistqueue"

type GridIndex struct {
    r int
    c int
}

func islandCheck(gridp *[][]byte, checkp *[][]bool, q *llq.Queue, g GridIndex) bool {
    grid := (*gridp)
    check := (*checkp)
    rows := len(grid) - 1
    cols := len(grid[0]) - 1
    if grid[g.r][g.c] == '0' { return false }
    check[g.r][g.c] = true

    if g.c > 0 && !check[g.r][g.c-1] && grid[g.r][g.c-1] == '1' {
        q.Enqueue(GridIndex{r:g.r, c:g.c-1,})
    }
    if g.c > 0 { check[g.r][g.c-1] = true }
    if g.c < cols && !check[g.r][g.c+1] && grid[g.r][g.c+1] == '1' {
        q.Enqueue(GridIndex{r:g.r, c:g.c+1,})
    }
    if g.c < cols { check[g.r][g.c+1] = true }
    if g.r > 0 && !check[g.r-1][g.c] && grid[g.r-1][g.c] == '1' {
        q.Enqueue(GridIndex{r:g.r-1, c:g.c,})
    }
    if g.r > 0 { check[g.r-1][g.c] = true }
    if g.r < rows && !check[g.r+1][g.c] && grid[g.r+1][g.c] == '1' {
        q.Enqueue(GridIndex{r:g.r+1, c:g.c,})
    }
     if g.r < rows { check[g.r+1][g.c] = true }
    return true
}

func numIslands(grid [][]byte) int {
    numIslands := 0
    if len(grid) == 0 { return numIslands }
    rows := len(grid) - 1
    cols := len(grid[0]) - 1
    checkGrid := make([][]bool, rows+1)
    for r := range checkGrid {
        checkGrid[r] = make([]bool, cols+1)
    }
    q := llq.New()
    for i := 0; i <= rows; i++ {
        for j := 0; j <= cols; j++ {
            if !checkGrid[i][j] && grid[i][j] == '1' {
                numIslands++
                q.Enqueue(GridIndex{r:i, c:j,})
                for !q.Empty() {
                    fEle, _ := q.Dequeue()
                    f := fEle.(GridIndex)
                    islandCheck(&grid, &checkGrid, q, f)
                }
            }
        }
    }
    return numIslands
}
