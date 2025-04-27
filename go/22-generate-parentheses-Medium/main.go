func generateParenthesis(n int) []string {
    res := []string{}

    var rec func([]byte, int, int)
    rec = func(prefix []byte, open, closed int) {
        if open == n && closed == n {
            res = append(res, string(prefix))
        }
        if open < n {
            rec(append(prefix, '('), open + 1, closed)
        }
        if closed < open {
            rec(append(prefix, ')'), open, closed + 1)
        }
    }

    rec(make([]byte, 0, 2 * n), 0, 0)
    return res
}