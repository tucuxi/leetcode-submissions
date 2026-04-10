func mirrorFrequency(s string) int {
    var c [26]int
    var d [10]int

    for _, b := range s {
        if b >= 'a' && b <= 'z' {
            c[b - 'a']++
        } else {
            d[b - '0']++
        }
    }

    res := 0
    
    for i, j := 0, 25; i <= j; i, j = i+1, j-1 {
        res += abs(c[i] - c[j])
    }
    for i, j := 0, 9; i <= j; i, j = i+1, j-1 {
        res += abs(d[i] - d[j])
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}