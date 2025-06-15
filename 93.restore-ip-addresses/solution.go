const parts = 4

func restoreIpAddresses(s string) []string {
    res := []string{}

    var parse func(string, []string)
    parse = func(s string, pre []string) {
        if len(pre) == parts && len(s) == 0 {
            res = append(res, strings.Join(pre, "."))
            return
        }
        if len(pre) == parts || len(s) == 0 {
            return
        }
        parse(s[1:], append(pre, s[:1]))
        if s[0] == '0' || len(s) <= 1 {
            return
        }
        parse(s[2:], append(pre, s[:2]))
        if len(s) <= 2 {
            return
        }
        v, _ := strconv.Atoi(s[:3])
        if v > 255 {
            return
        }
        parse(s[3:], append(pre, s[:3]))
    }
    
    parse(s, make([]string, 0, parts))
    return res
}
