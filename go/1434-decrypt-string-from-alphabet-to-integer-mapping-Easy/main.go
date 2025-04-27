func freqAlphabets(s string) string {
    nums := []byte{}
    for i := len(s) - 1; i >= 0; {
        if s[i] == '#' {
            n := 10 * (s[i - 2] - '0') + s[i - 1] - '0'
            nums = append(nums, n)
            i -= 3
        } else {
            n := s[i] - '0'
            nums = append(nums, n)
            i--
        }
    }
    var b strings.Builder
    for i := len(nums) - 1; i >= 0; i-- {
        b.WriteByte('a' + nums[i] - 1)
    }
    return b.String()
}