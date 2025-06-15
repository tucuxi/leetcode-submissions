func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    equiv := func() map[byte]byte {
        equiv := make(map[byte]byte)
        for c := byte('a'); c <= 'z'; c++ {
            equiv[c] = c
        }
        return equiv
    }()

    find := func(x byte) byte {
        for equiv[x] != x {
            x, equiv[x] = equiv[x], equiv[equiv[x]]
        }
        return x
    }

    for i := range s1 {
        c1 := find(s1[i])
        c2 := find(s2[i])
        if c1 < c2 {
            equiv[c2] = c1
        } else {
            equiv[c1] = c2
        }
    }

    var sb strings.Builder
    for i := range baseStr {
        c := find(baseStr[i])
        sb.WriteByte(c)
    }
    
    return sb.String()
}