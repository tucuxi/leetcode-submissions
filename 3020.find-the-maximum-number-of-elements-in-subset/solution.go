func maximumLength(nums []int) int {
    c := make(map[int]int)
    
    for _, num := range nums {
        c[num]++
    }

    res := c[1]
    delete(c, 1)
    
    if res%2 == 0 {
        res--
    }

    for num := range c {
        x := num
        l := 0
        
        for c[x] > 1 {
            l += 2
            x *= x
        }
        if c[x] > 0 {
            res = max(res, l+1)
        } else {
            res = max(res, l-1)
        }
    }

    return res
}