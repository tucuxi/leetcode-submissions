type bits struct {
    count [32]int
}

func (b bits) add(n int) bits {
    res := b
    for i := range b.count {
        if n & (1 << i) != 0 {
            res.count[i]++
        }
    }
    return res
}

func (b bits) sub(n int) bits {
    res := b
    for i := range res.count {
        if n & (1 << i) != 0 {
            res.count[i]--
        }
    }
    return res
}

func (b bits) value() int {
    v := 0
    for i := range b.count {
        if b.count[i] != 0 {
            v |= 1 << i
        }
    }
    return v
}

func minimumSubarrayLength(nums []int, k int) int {
    res := math.MaxInt
    p := 0
    w := 0
    b := bits{}
    for q := range nums {
        w |= nums[q]
        b = b.add(nums[q])
        if w >= k { 
            for p < q && w >= k {
                c := b.sub(nums[p])
                v := c.value()
                if v < k {
                    break
                }
                b = c
                w = v
                p++
            }
            res = min(res, q - p + 1)
        }
    }
    if res == math.MaxInt {
        return -1
    }
    return res
}