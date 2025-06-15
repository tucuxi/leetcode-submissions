func romanToInt(s string) int {
    d := [...]string{"IV", "IX", "XL", "XC", "CD", "CM", "I", "V", "X", "L", "C", "D", "M"}
    v := [...]int{4, 9, 40, 90, 400, 900, 1, 5, 10, 50, 100, 500, 1000}
    res := 0
    outer: for len(s) > 0 {
        for i, p := range d {
            if strings.Index(s, p) == 0 {
                res += v[i]
                s = s[len(p):]
                continue outer
            }
        }
        return 0
    }
    return res
}