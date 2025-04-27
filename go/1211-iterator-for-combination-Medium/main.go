type CombinationIterator struct {
    characters string
    combinations []uint16
    next int
}


func Constructor(characters string, combinationLength int) CombinationIterator {
    var c []uint16
    var enumerate func(int, uint16, int)
    
    enumerate = func(index int, acc uint16, taken int) {
        if taken == combinationLength {
            c = append(c, acc)
            return
        }
        if taken > combinationLength || index == len(characters) {
            return
        }
        enumerate(index + 1, acc | uint16(1 << index), taken + 1)
        enumerate(index +1, acc, taken)
    }
    
    enumerate(0, 0, 0)
    
    return CombinationIterator{
        characters:   characters,
        combinations: c,
        next:         0,
    }
    
}


func (this *CombinationIterator) Next() string {
    var b strings.Builder
    c := this.combinations[this.next]
    for i := range this.characters {
        if c & uint16(1 << i) != 0 {
            b.WriteByte(this.characters[i])
        }
    }
    this.next++
    return b.String()
}


func (this *CombinationIterator) HasNext() bool {
    return this.next < len(this.combinations)
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */