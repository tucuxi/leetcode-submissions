func minimumDeleteSum(s1 string, s2 string) int {
    memo := make([][]int, len(s1))
    for i := range memo {
        memo[i] = make([]int, len(s2))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var f func(int, int) int
    f = func(i1, i2 int) int {
        if i1 == len(s1) && i2 == len(s2) {
            return 0
        }
        if i1 == len(s1) {
            return sum(s2[i2:])
        }
        if i2 == len(s2) {
            return sum(s1[i1:])
        }
        if memo[i1][i2] != -1 {
            return memo[i1][i2]
        }
        if s1[i1] == s2[i2] {
            memo[i1][i2] = f(i1 + 1, i2 + 1)
        } else {
            memo[i1][i2] = min(f(i1 + 1, i2) + int(s1[i1]), f(i1, i2 + 1) + int(s2[i2]))
        }
        return memo[i1][i2]
    }

    return f(0, 0)
}

func sum(s string) int {
    res := 0
    for _, ch := range s {
        res += int(ch)
    }
    return res
}