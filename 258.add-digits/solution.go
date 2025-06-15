func addDigits(num int) int {
    for num > 9 {
        n := 0
        for m := num; m > 0; m /= 10 {
            n += m % 10
        }
        num = n
    }
    return num
}