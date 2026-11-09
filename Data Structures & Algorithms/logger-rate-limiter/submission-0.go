type Logger struct {
    m map[string]int
}


func Constructor() Logger {
    return Logger{
        m: make(map[string]int),
    }
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    v, e := this.m[message]
    if e && timestamp < (v+10) {
        return false
    }
    this.m[message] = timestamp
    return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
