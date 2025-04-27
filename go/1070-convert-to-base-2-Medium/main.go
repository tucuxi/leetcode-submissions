func baseNeg2(n int) string {
    if n == 0 {
        return "0"
    }

    var r []byte

    for num := n; num != 0; num = -(num >> 1) {
        r = append([]byte{byte(num & 1) + '0'}, r...)
    }

    return string(r)
}
