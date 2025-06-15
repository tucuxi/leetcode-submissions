type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    return Solution{nums}
}


func (this *Solution) Reset() []int {
    return this.nums
}


func (this *Solution) Shuffle() []int {
    res := make([]int, len(this.nums))
    copy(res, this.nums)
    for i := len(res) - 1; i >= 1; i-- {
        j := rand.Intn(i + 1)
        res[i], res[j] = res[j], res[i]
    }
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */