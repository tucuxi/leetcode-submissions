func countKthRoots(l int, r int, k int) int {
    if r == 0 {
        return 1
    }

    if k == 1 {
        return r-l+1
    }

    lr, rr := 0, 0

    for x := 1;; x++ {
        y := 1
        for range k {
            y *= x
        }
        if lr == 0 && y >= l {
            lr = x
        }
        if y > r {
            rr = x-1
            break
        }
    }

    if l == 0 {
        return rr+1
    }
    return rr-lr+1
}