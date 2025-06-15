type RandomizedSet struct {
    set map[int]bool
    elems []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{set: map[int]bool{}, elems: []int{}}    
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.set[val]; !exists {
        this.set[val] = true
        this.elems = append(this.elems, val)
        return true
    }
    return false  
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, exists := this.set[val]; exists {
        delete(this.set, val)
        for i, v := range this.elems {
            if v == val {
                this.elems = append(this.elems[:i], this.elems[i+1:]...)
                return true
            }
        }
    }
    return false
}


func (this *RandomizedSet) GetRandom() int {
    r := rand.Intn(len(this.elems))
    return this.elems[r]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */