func findPermutationDifference(s string, t string) int {
    var tIndex [26]int
    
    for i := range t {
        tIndex[t[i] - 'a'] = i
    }

    res := 0

    for i := range s {
        difference := abs(i - tIndex[s[i] - 'a'])
        res += difference
    }

    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}