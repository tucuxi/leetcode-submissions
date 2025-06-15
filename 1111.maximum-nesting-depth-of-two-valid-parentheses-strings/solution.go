func maxDepthAfterSplit(seq string) []int {
    var res []int
    d := 0
    for _, r := range seq {
        if r == '(' {
            d++
        }
        res = append(res, d % 2)
        if r == ')' {
            d--
        }
    }
    return res
}