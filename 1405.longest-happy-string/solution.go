type frequency struct {
    b byte
    n int
}

func longestDiverseString(a int, b int, c int) string {
    var res []byte
    h := []frequency{frequency{'a', a}, frequency{'b', b}, frequency{'c', c}}
    
    for i := 0;; i++ {
        sort.Slice(h, func(i, j int) bool { return h[i].n < h[j].n })
        if h[2].n == 0 {
            break
        }
        if i >= 2 && res[i-1] == h[2].b && res[i-2] == h[2].b {
            if h[1].n == 0 {
                break
            }
            h[1].n--
            res = append(res, h[1].b)
        } else {
            h[2].n--
            res = append(res, h[2].b)
        }
    }
    return string(res)
}