func validStrings(n int) []string {
    var res []string
    var rec func(string)
    
    rec = func(pre string) {
        if len(pre) == n {
            res = append(res, pre)
            return
        }
        rec(pre + "1")
        if len(pre) == 0 || pre[len(pre) - 1] == '1' {
            rec(pre + "0")
        }
    }

    rec("")
    return res
}