func minimizeXor(num1 int, num2 int) int {
    x := 0
    b := countBits(num2)
    for m := 1 << 32; b > 0 && m != 0; m >>= 1 {
        if num1 & m != 0 {
            x |= m
            b--
        }
    }
    for m := 1; b > 0; m <<= 1 {
        if x & m == 0 {
            x |= m
            b--
        }
    }
    return x
}

func countBits(num int) int {
    b := 0
    for num != 0 {
        num &= num - 1
        b++
    }
    return b
}