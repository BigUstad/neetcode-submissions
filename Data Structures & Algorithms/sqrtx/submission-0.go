func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    l := 1
    r := x
    xl := uint64(x)
    // Binary search for m
    // Where m*m <= x & (m+1)*(m+1) > x
    for l <= r {
        m := l + (r - l)/2
        sq := uint64(m * m)
        nextSq := uint64((m+1)*(m+1))
        // fmt.Println(m)
        if sq <= xl && nextSq > xl {
            return m
        }
        if sq > xl {
            r = m - 1
        } else {
            l = m + 1
        }
    }
    return -1
}
