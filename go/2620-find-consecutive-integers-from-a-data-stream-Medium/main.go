type DataStream struct {
    value int
    length int
    run int
}


func Constructor(value int, k int) DataStream {
    return DataStream{value, k, 0}
}


func (this *DataStream) Consec(num int) bool {
    if num == this.value {
        this.run++
    } else {
        this.run = 0
    }
    return this.run >= this.length
}


/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */