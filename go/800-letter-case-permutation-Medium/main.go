func letterCasePermutation(s string) []string {
    res := []string{}
    buf := []byte(s)

    var rec func(int)
    rec = func(index int) {
        if index == len(buf) {
            res = append(res, string(buf))
            return
        }
        rec(index + 1)
        if buf[index] >= 'A' {
            buf[index] ^= 0x20
            rec(index + 1)
            buf[index] ^= 0x20
        }
    }

    rec(0)
    return res
}