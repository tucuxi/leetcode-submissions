type NumArray struct {
    bit []int
}


func Constructor(nums []int) NumArray {
    bit := make([]int, len(nums) + 1)
    for i := 1; i < len(bit); i++ {
        bit[i] += nums[i - 1]
        j := i + (-i & i)
        if j < len(bit) {
            bit[j] += bit[i]
        }
    }
    return NumArray{bit}
}


func (this *NumArray) Update(index int, val int)  {
    d := val - this.SumRange(index, index)
    for i := index + 1; i < len(this.bit); i += -i & i {
        this.bit[i] += d
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    sum := 0 
    i := left 
    j := right + 1
    for ; j > i; j -= -j & j {
        sum += this.bit[j]
    }
    for ; i > j; i -= -i & i {
        sum -= this.bit[i]
    }
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */