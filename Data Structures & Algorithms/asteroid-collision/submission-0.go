func asteroidCollision(asteroids []int) []int {
    var stck []int
    // var res []int
    l := len(asteroids)
    // Stack maintains surviving asteroids till "current" i.e index "i"
    // If they survive till the end they will be added to result
    // Solution from leetcode
    for i := 0; i < l; i++ {
        addToStack := true
        // Stack has "positive" asteroids & the current asteroid is "negative"
        for len(stck) > 0 && stck[len(stck)-1] > 0 && asteroids[i] < 0 {
            // fmt.Println("1. ", stck)
            top := stck[len(stck) - 1]
            cmpTop := top
            cmpCur := -asteroids[i]
            if top < 0 { cmpTop = -top }
            if cmpTop < cmpCur {
                // If stack top is smaller than current. stack top asteroid will get destroyed
                // So pop it
                stck = stck[:len(stck)-1]
                continue
            } else if cmpTop == cmpCur {
                // If stack top is equal to current. Both will be destroyed.
                // So pop it. Also set addToStack to false as current also got destroyed.
                stck = stck[:len(stck)-1]
                // NOTE: This code is slightly redundant but left as-is for clarity
                addToStack = false
                break
            }
            // cmpTop > cmpCur. So asteroids[i] will be destroyed
            // flag set to false as the asteroids[i] will not be added to stack
            addToStack = false
            break
        }
        if addToStack {
            stck = append(stck, asteroids[i])
        }
        // fmt.Println("2. ", stck)
    }
    return stck
}
