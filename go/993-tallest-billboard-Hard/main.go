type states map[[2]int]struct{}

func tallestBillboard(rods []int) int {

    gen := func(left, right int) map[int]int {
        all := make(states)
        all[[2]int{0, 0}] = struct{}{}
        for i := left; i < right; i++ {
            r := rods[i]
            next := make(states)
            for p := range all {
                next[[2]int{p[0] + r, p[1]}] = struct{}{}
                next[[2]int{p[0], p[1] + r}] = struct{}{}
            }
            for p := range next {
                all[p] = struct{}{}
            }
        }
        res := make(map[int]int)
        for p := range all {
            l, r := p[0], p[1]
            res[l - r] = max(l, res[l - r])
        }
        return res
    }
    
    l := gen(0, len(rods) / 2)
    r := gen(len(rods) / 2, len(rods))
    
    res := 0
    for lk, lv := range l {
        if rv, exists := r[-lk]; exists {
            res = max(res, lv + rv)
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
