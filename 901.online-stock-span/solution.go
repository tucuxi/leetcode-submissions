type StockSpanner struct {
    quotes []int16
}


func Constructor() StockSpanner {
    return StockSpanner{[]int16{}}
}


func (this *StockSpanner) Next(price int) int {
    var p = int16(price)
    this.quotes = append(this.quotes, p)
    for i := len(this.quotes) - 2; i >= 0; i-- {
        if this.quotes[i] > p {
            return len(this.quotes) - i - 1
        }
    }
    return len(this.quotes)
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */