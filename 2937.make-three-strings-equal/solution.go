func findMinimumOperations(s1 string, s2 string, s3 string) int {
    n := min(len(s1), len(s2), len(s3))
    i := 0
    for i < n && s1[i] == s2[i] && s1[i] == s3[i] {
        i++
    }
    if i == 0 {
        return -1
    }
    return len(s1) + len(s2) + len(s3) - 3 * i
}

func min(a ...int) int {
    res := a[0]
    for _, n := range a {
        if n < res {
            res = n
        }
    }
    return res
}