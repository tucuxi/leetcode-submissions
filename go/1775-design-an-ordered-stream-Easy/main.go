type OrderedStream struct {
    values []string
    ptr int
}


func Constructor(n int) OrderedStream {
    return OrderedStream{make([]string, n), 0}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.values[idKey - 1] = value
    p := this.ptr
    for this.ptr < len(this.values) && this.values[this.ptr] != "" {
        this.ptr++
    }
    return this.values[p:this.ptr]
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */