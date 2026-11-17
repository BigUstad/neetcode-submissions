func lemonadeChange(bills []int) bool {
    if len(bills) == 0 || bills[0] == 10 || bills[0] == 20 {
        return false
    }
    fives := 0
    tens := 0
    for _, b := range bills {
        // fmt.Println(b)
        switch b {
            case 5:
                fives++
            case 10:
                tens++
                fives--
            case 20:
                if tens > 0 {
                    tens--
                    fives--
                } else {
                    fives -= 3
                }
            default:
                return false
        }
        if fives < 0 || tens < 0 {
            return false
        }
        // fmt.Println(fives, tens)
   }
    return true
}
