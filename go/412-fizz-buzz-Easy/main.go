func fizzBuzz(n int) []string {
    res := make([]string, 0, n)
    for i := 1; i <= n; i++ {
        var s string
        if i % 3 == 0 {
            s = "Fizz"
        }
        if i % 5 == 0 {
            s += "Buzz"
        }
        if len(s) == 0 {
            s = strconv.Itoa(i)
        }
        res = append(res, s)
    }
    return res
}