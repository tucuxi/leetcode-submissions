func largestGoodInteger(num string) string {
    var (
        prev, max rune
        run = 1
        res = ""
    )
    
    for i, r := range num {
        if r == prev && r > max {
            run++
            if run == 3 {
                max = r
                res = num[i-2:i+1]
            }
        } else {
            prev = r
            run = 1
        }
    }
    return res
}