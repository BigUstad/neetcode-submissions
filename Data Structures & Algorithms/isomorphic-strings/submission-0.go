func isIsomorphic(s string, t string) bool {
    // Let's create the mappings of string s char to string t char if it doesn't exist.
    // Let's recreate the string t with the mapping. 
    // If the same string isn't recreated. Then there is a mismatch.
    // Create a string transformation that proves or disproves the requires condition
    lns := len(s)
    lnt := len(t)
    mst := make(map[rune]rune)
    // TO check no 2 characters in s map to same character in t
    mts := make(map[rune]rune)
    if lns != lnt {
        return false
    }
    // Every char in t will be mapped to char same index in s
    for i := 0; i < lns; i++ {
        rs := rune(s[i])
        rt := rune(t[i])
        ert, ok1 := mst[rs]
        ers, ok2 := mts[rt]
        if ok1 && ert != rt {
            return false
        } else if ok2 && ers != rs {
            return false
        }
        mst[rs] = rt
        mts[rt] = rs
    }

    return true
}
