func hasAllCodes(s string, k int) bool {
    if len(s) < k {
        return false
    }

    var (
        c     uint32
        m     uint32 = (1 << k) - 1
        codes        = make([]bool, 1 << k)
    )
        
    for i := range k {
        c = (c << 1) | uint32(s[i] - '0')
    }

    codes[c] = true

    for i := k; i < len(s); i++ {
        c = (c << 1) & m | uint32(s[i] - '0')
        codes[c] = true
    }

    for _, b := range codes {
        if !b {
            return false
        }
    }

    return true
}