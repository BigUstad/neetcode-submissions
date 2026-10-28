func checkValidString(s string) bool {
    // TODO: The stack solution is more interesting & has a small learning.
    // THis is an attempt of the Greedy solution
    // 1. Initialize two counters: leftMin leftMax
    leftMin, leftMax := 0, 0
    // 2. Traverse the string character by character:
    // 3. For each character c:
    //      if c is '(' increment both leftmin & leftmax
    //      if c is ')' decrement both leftmin & leftmax
    //      if c is '*', then flexible treatment is needed.
    //      you can decrement leftmin (consider it ')')
    //      you can increment leftmax (condider it '(' )
    // 4. if leftMax < 0 (too many closing parantheses)
    // 5. if leftmin < 0 reset leftmin to 0 (extra closings treated as empty)
    // 6. Outside the loop, if leftmin = 0 then return true. Else false
    for _, c := range s {
        if c == '(' {
            leftMin++
            leftMax++
        } else if c == ')' {
            leftMin--
            leftMax--
        } else {
            // This is for '*'
            // Let's decrement leftmin
            // And increment leftmax
            leftMin--
            leftMax++
        }
        if leftMax < 0 {
            // ')' has clearly exceeded
            // '(' even if they come afterwards, won't satisfy
            return false
        }
        if leftMin < 0 {
            // A case for '*'s exceeding because leftMax is still above 0
            // Let's treat * as empty
            leftMin = 0
        }
    }
    return leftMin == 0
}
