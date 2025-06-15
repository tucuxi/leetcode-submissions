type RLEIterator struct {
    encoding []int
    i1 int
    i2 int
}


func Constructor(encoding []int) RLEIterator {
    var e []int
    for i := 0; i < len(encoding); i += 2 {
        if encoding[i] > 0 {
            e = append(e, encoding[i], encoding[i+1])
        }
    }
    return RLEIterator{e, 0, 0}
}


func (this *RLEIterator) Next(n int) int {
    for n > 0 && this.i1 < len(this.encoding) {
        res := this.encoding[this.i1 + 1]
        l := this.encoding[this.i1] - this.i2
        if n == l {
            this.i1 += 2
            this.i2 = 0
            return res
        }
        if n < l {
            this.i2 += n
            return res
        }
        this.i1 += 2
        this.i2 = 0
        n -= l
    }
    return -1
}