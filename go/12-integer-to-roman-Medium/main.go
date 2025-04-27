func intToRoman(num int) string {
    var res strings.Builder
    digits := []byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
    add(num / 1000, 'M', &res)
    place(num % 1000 / 100, digits[:3], &res)
    place(num % 100 / 10, digits[2:5], &res)
    place(num % 10, digits[4:], &res)
    return res.String()
}

func place(n int, digits []byte, res *strings.Builder) {
    switch n {
        case 0, 1, 2, 3:
            add(n, digits[2], res)
        case 4:
            add(1, digits[2], res)
            add(1, digits[1], res)
        case 5, 6, 7, 8:
            add(1, digits[1], res)
            add(n - 5, digits[2], res)
        case 9:
            add(1, digits[2], res)
            add(1, digits[0], res)
    }
}

func add(n int, char byte, sb *strings.Builder) {
    for i := 0; i < n; i++ {
        sb.WriteByte(char)
    }
}