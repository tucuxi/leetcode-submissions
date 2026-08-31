func isPalindromic(s string) bool {
    i := 0
    j := len(s) - 1

    for i <= j {
        if !palindromic(s[i], s[j]) {
            return false
        }
        i++
        j--
    }
    return true
}

func palindromic(a, b byte) bool {
    var ma, mb byte = 128, 1

    for ma > 0 {
        if (a & ma == 0) != (b & mb == 0) {
            return false
        }
        ma >>= 1
        mb <<= 1
    }
    return true
}