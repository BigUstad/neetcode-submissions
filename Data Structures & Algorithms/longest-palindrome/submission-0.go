func longestPalindrome(s string) int {
    res := 0
    oneCount := 0
    // 1 to 52
    cm := make([]uint16, 53)
    for _, c := range s {
        i := c - 'A'
        if !unicode.IsUpper(c) {
            // 'a' - 97, 'A' - 65
            i = c - 'A' - 5
        }
        // fmt.Fprintf(os.Stdout, "%c...%d\n", c, i)
        cm[i]++
    }
    // fmt.Println(cm)
    for _, v := range cm {
        
        if v%2 == 0 {
            res += int(v)
        } else {
            res += int(v - 1)
            oneCount++
        }
    }
    if oneCount >= 1 { res++ }
    return res
}
