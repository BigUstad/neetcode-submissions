func mergeAlternately(word1 string, word2 string) string {
    i := 0
    j := 0
    lnw1 := len(word1)
    lnw2 := len(word2)
    var b strings.Builder
    for i < lnw1 && j < lnw2 {
        b.WriteByte(word1[i])
        b.WriteByte(word2[i])
        i++
        j++
    }
    if i < lnw1 {
        b.WriteString(word1[i:])
    }
    if j < lnw2 {
        b.WriteString(word2[j:])
    }
    return b.String()
}
