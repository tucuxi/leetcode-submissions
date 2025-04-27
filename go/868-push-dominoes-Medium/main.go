func pushDominoes(dominoes string) string {
    s, n := dominoes, len(dominoes)
    for {
        b := []byte(s)
        changed := false
        for i := 0; i < n; i++ {
            if s[i] == '.' {
                if i > 0 && s[i-1] == 'R' && (i == n-1 || s[i+1] != 'L') {
                    b[i] = 'R'
                    changed = true
                } else if (i == 0 || s[i-1] != 'R') && i < n-1 && s[i+1] == 'L' {
                    b[i] = 'L'
                    changed = true
                }
            }
        }
        if !changed {
            return s
        }
        s = string(b)
    }
}