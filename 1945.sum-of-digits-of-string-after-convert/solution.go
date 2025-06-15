func getLucky(s string, k int) int {
    res := convertAndSum(s)
    for range k-1 {
        res = transform(res)
    }    
    return res
}

func convertAndSum(s string) int {
    res := 0
    for _, ch := range s {
        if n := int(ch - 'a') + 1; n >= 10 {
            res += n / 10 + n % 10
        } else {
            res += n
        }
    }
    return res
}

func transform(num int) int {
    res := 0
    for ; num > 0; num /= 10 {
        res += num % 10
    }
    return res
}