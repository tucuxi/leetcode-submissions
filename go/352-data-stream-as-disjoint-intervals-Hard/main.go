type SummaryRanges struct {
    intervals [][]int
}


func Constructor() SummaryRanges {
    return SummaryRanges{} 
}


func (this *SummaryRanges) AddNum(value int)  {
    k := sort.Search(len(this.intervals), func(i int) bool {
        return this.intervals[i][0] >= value
    })
    if k < len(this.intervals) && this.intervals[k][0] == value || k > 0 && this.intervals[k - 1][1] >= value {
        return
    }
    if k > 0 && this.intervals[k - 1][1] == value - 1 {
        if k < len(this.intervals) && this.intervals[k][0] == value + 1 {
            this.intervals[k - 1][1] = this.intervals[k][1]
            this.intervals = append(this.intervals[:k], this.intervals[k + 1:]...)
        } else {
            this.intervals[k- 1][1] = value
        }
        return
    }
    if k < len(this.intervals) && this.intervals[k][0] == value + 1 {
        this.intervals[k][0] = value
    } else {
        this.intervals = append(this.intervals[:k], append([][]int{{value, value}}, this.intervals[k:]...)...)
    }
}


func (this *SummaryRanges) GetIntervals() [][]int {
    return this.intervals 
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */