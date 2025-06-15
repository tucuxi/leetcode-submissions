func findEvenNumbers(digits []int) []int {
    all := make(map[int]bool)
    for i := range digits {
        if digits[i] == 0 {
            continue
        }
        for j := range digits{
            if i == j {
                continue
            }
            for k := range digits {
                if i == k || j == k || digits[k] % 2 != 0 {
                    continue
                }
                n := 100 * digits[i] + 10 * digits[j] + digits[k]
                all[n] = true
            }
        }
    }
    var res []int
    for num := range all {
        res = append(res, num)
    }
    slices.Sort(res)
    return res
}
