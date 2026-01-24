type Solution struct {
    prefix []int
    limit  int
}


func Constructor(w []int) Solution {
    prefix := make([]int, 0, len(w))
    sum := 0
    for _, weight := range w {
        sum += weight
        prefix = append(prefix, sum)
    }
    return Solution{
        prefix: prefix,
        limit:  sum,
    }
}


func (this *Solution) PickIndex() int {
    r := rand.Intn(this.limit) + 1
    i := sort.Search(len(this.prefix), func(i int) bool { return this.prefix[i] >= r })
    return i
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */