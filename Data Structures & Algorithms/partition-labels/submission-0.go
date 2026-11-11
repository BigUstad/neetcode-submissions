func partitionLabels(s string) []int {
    // map of char to last index
    // O(n) time. O(1) as there are only 26 characters.
    m := make(map[rune]int)
    for i, c := range s {
        m[rune(c)] = i
    }
    var res []int
    size, end := 0, 0
    // Iterate over string again, char by char
    for i, c := range s {
        r := rune(c)
        size++
        end = max(end, m[r])
        if i == end {
            res = append(res, size)
            size = 0
        }
    }
    return res
}
