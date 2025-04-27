func findEvenNumbers(digits []int) []int {
    all := map[int]bool{}
    cnt := len(digits)
    for i := 0; i < cnt; i++ {
        for j := 0; j < cnt; j++ {
            if i == j {
                continue
            }
            for k := 0; k < cnt; k++ {
                if i == k || j == k {
                    continue
                }
                n := 100 * digits[i] + 10 * digits[j] + digits[k]
                if n >= 100 && n % 2 == 0 {
                    all[n] = true
                }
            }
        }
    }
    res := []int{}
    for num := range all {
        res = append(res, num)
    }
    sort.Ints(res)
    return res
}
