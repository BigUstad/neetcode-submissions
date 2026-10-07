func decodeString(s string) string {
    numstack := list.New()
    stringstack := list.New()
    var cur string
    k := 0
    for _, c := range s {
        if unicode.IsDigit(c) {
            ci := int((c - '0'))
            k = k*10 + ci
        } else if c == '[' {
            if k > 0 {
                numstack.PushFront(k)
            }
            stringstack.PushFront(cur)
            k = 0
            cur = ""
        } else if c == ']' {
            ne := numstack.Remove(numstack.Front())
            ni := ne.(int)
            decoded := ""
            if stringstack.Len() > 0 {
                se := stringstack.Remove(stringstack.Front())
                decoded = se.(string)
            }
            for i := ni; i > 0; i-- {
                decoded = decoded + cur
            }
            cur = decoded
        } else {
            cur = cur + string(c)
        }
    }
    return cur
}
