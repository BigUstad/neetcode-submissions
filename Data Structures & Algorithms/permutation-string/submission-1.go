import ("maps")

func checkInclusion(s1 string, s2 string) bool {
    l1 := len(s1)
    l2 := len(s2)
    if l2 < l1 { return false }
    s1map, s2map := make(map[rune]int), make(map[rune]int)
    for i, c := range s1 {
        s1map[c]++
        s2map[rune(s2[i])]++
    }
    if maps.Equal(s1map, s2map) {
        return true
    }
    i := 1
    for j := l1; j < l2; j++ {
        // fmt.Println(s1map)
        // fmt.Println(s2map)
        // fmt.Println("---------")
        r := rune(s2[i-1])
        re := rune(s2[j])
        // fmt.Print(r)
        // fmt.Print(",")
        // fmt.Println(re)
        // fmt.Println("==================")
        if s2map[r] == 1 {
            delete(s2map, r)
        } else if s2map[r] > 1 {
            s2map[r]--
        }
        s2map[re]++
        i++
        if maps.Equal(s1map, s2map) {
            return true
        }
    }
    return false
}
