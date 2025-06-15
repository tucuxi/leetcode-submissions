func minMaxDifference(num int) int {
    s := strconv.Itoa(num)
    max, _ := strconv.Atoi(replaceFirstNotEqual(s, '9'))
    min, _ := strconv.Atoi(replaceFirstNotEqual(s, '0'))
    return max - min
}

func replaceFirstNotEqual(s string, r byte) string {
    var sb strings.Builder
    i := 0

    for i < len(s) && s[i] == r {
        sb.WriteByte(r)
        i++
    }
    if i < len(s) {
        found := s[i]
        for j := i; j < len(s); j++ {
            if s[j] == found {
                sb.WriteByte(r)
            } else {
                sb.WriteByte(s[j])
            }
        }
    }
    return sb.String()
}