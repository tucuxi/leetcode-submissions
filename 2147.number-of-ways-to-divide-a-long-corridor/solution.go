func numberOfWays(corridor string) int {
    const mod = 1e9+7
    res, s, p := 1, 0, 0
    
    for _, c := range corridor {
        if c == 'S' {
            if s == 2 {
                res = res * (p + 1) % mod
                s, p = 0, 0
            }
            s++
        } else if s == 2 {
            p++
        }
    }
    
    if s < 2 {
        return 0
    }
    return res
}