func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    equiv := map[byte]byte{}
    for c := byte('a'); c <= 'z'; c++ {
        equiv[c] = c
    }
    for i := range s1 {
        c1 := find(equiv, s1[i])
        c2 := find(equiv, s2[i])
        if c1 < c2 {
            equiv[c2] = c1
        } else {
            equiv[c1] = c2
        }
    }
    var sb strings.Builder
    for i := range baseStr {
        c := find(equiv, baseStr[i])
        sb.WriteByte(c)
    }
    return sb.String()
}

func find(uf map[byte]byte, x byte) byte {
    for uf[x] != x {
        x, uf[x] = uf[x], uf[uf[x]]
    }
    return x
}