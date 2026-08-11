// import "github.com/emirpasic/gods/stacks/arraystack"

const (
    open_param = '('
    close_param = ')'
    open_sq = '['
    close_sq = ']'
    open_fl = '{'
    close_fl = '}'
)

func isValid(s string) bool {
    if (len(s) == 0) {
        return true
    }
    if (len(s) % 2) != 0 {
        return false
    }
    stck := list.New()
    for _, c := range s {
        is_open :=  (c == open_param ||
                     c == open_sq ||
                     c == open_fl)
        if is_open {
            stck.PushFront(c)
            continue
        }
        if stck.Len() == 0 {
            return false
        }
		top := stck.Front().Value
		stck.Remove(stck.Front())
        is_mismatched := (top == open_param && c != close_param ||
                          top == open_sq && c != close_sq ||
                          top == open_fl && c != close_fl)
        if is_mismatched {
            return false
        }
    }
    return (stck.Len() == 0)
}
