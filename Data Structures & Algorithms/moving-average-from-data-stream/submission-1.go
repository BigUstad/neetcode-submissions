type MovingAverage struct {
    istrm []int
    size float64
    curSum float64
}

func Constructor(size int) *MovingAverage {
    return &MovingAverage{
        istrm: nil,
        size: float64(size),
        curSum: float64(0),
    }
}

func (this *MovingAverage) Next(val int) float64 {
    this.istrm = append(this.istrm, val)
    curSize := len(this.istrm)
    if curSize <= int(this.size) {
        this.curSum += float64(val)
        avg := this.curSum / float64(curSize)
        // fmt.Println(avg, " size= ", float64(len(this.istrm)))
        return avg
    }
    // Remove the beginning of the istrm slice
    // As soon as length of slice is greater than existing slice
    first := this.istrm[0]
    this.istrm = this.istrm[1:]
    // fmt.Println("oldsum = ", this.curSum, " new size= ", float64(len(this.istrm)))
    this.curSum = this.curSum - float64(first) + float64(val)
    avg := this.curSum / this.size
    // fmt.Println("newsum = ", this.curSum, " avg= ", avg, " size= ", float64(len(this.istrm)))
    return avg
}
