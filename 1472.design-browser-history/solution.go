type BrowserHistory struct {
    history []string
    current int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        history: []string{homepage},
        current: 0,
    }
}


func (this *BrowserHistory) Visit(url string)  {
    this.current++
    if this.current == len(this.history) {
        this.history = append(this.history, url)
    } else {
        this.history[this.current] = url
        this.history = this.history[:this.current + 1]
    }
}


func (this *BrowserHistory) Back(steps int) string {
    index := this.current - steps
    if index < 0 {
        index = 0
    }
    this.current = index
    return this.history[index]
}


func (this *BrowserHistory) Forward(steps int) string {
    index := this.current + steps
    if index >= len(this.history) {
        index = len(this.history) - 1
    }
    this.current = index
    return this.history[index]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */