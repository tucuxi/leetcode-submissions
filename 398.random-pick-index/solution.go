type Solution struct {
    tab map[int][]int
}


func Constructor(nums []int) Solution {
    tab := map[int][]int{}
    for i, n := range nums {
        tab[n] = append(tab[n], i)
    }
    return Solution{tab}
}


func (this *Solution) Pick(target int) int {
    idx := this.tab[target]
    return idx[rand.Intn(len(idx))]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */