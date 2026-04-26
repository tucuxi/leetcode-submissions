func validDigit(n int, x int) bool {
    s := strconv.Itoa(n)
    c := byte('0' + x)
    return strings.IndexByte(s, c) > 0
}