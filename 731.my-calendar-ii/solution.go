type MyCalendarTwo struct {
    time []int
    change []int
}

func Constructor() MyCalendarTwo {
    return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
    if this.count(start, end) == 2 {
        return false
    }
    this.add(start, 1)
    this.add(end, -1)
    return true
}

func (this *MyCalendarTwo) count(start, end int) int {
    cur := 0
    i := 0
    for ; i < len(this.change) && this.time[i] <= start; i++ {
        cur += this.change[i]
    }
    max := cur
    for ; i < len(this.change) && this.time[i] < end; i++ {
        cur += this.change[i]
        if cur > max {
            max = cur
        }
    }
    return max
}

func (this *MyCalendarTwo) add(time int, change int) {
    i := sort.Search(len(this.time), func(i int) bool {
        return this.time[i] > time
    })
    if i > 0 && this.time[i-1] == time {
        this.change[i-1] += change
    } else {
        this.time = append(this.time, 0)
        copy(this.time[i+1:], this.time[i:])
        this.time[i] = time
        this.change = append(this.change, 0)
        copy(this.change[i+1:], this.change[i:])
        this.change[i] = change
    }
}