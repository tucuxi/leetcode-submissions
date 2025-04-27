func equationsPossible(equations []string) bool {
    uf := make([]byte, 26)
    for i := range uf {
        uf[i] = byte(i)
    }
    for _, eq := range equations {
        if eq[1] == '=' {
            union(uf, eq[0] - 'a', eq[3] - 'a')
        }
    }
    for _, eq := range equations {
        if eq[1] == '!' {
            if find(uf, eq[0] - 'a') == find(uf, eq[3] - 'a') {
                return false
            }
        }
    }
    return true
}

func union(uf []byte, a, b byte) {
    uf[find(uf, a)] = uf[find(uf, b)]    
}

func find(uf []byte, x byte) byte {
    if uf[x] != x {
        uf[x] = find(uf, uf[x])
    }
    return uf[x]    
}