func diffWaysToCompute(expression string) []int {
    var res []int
    num := 0
    for i := range expression {
        switch expression[i] {
        case '+':
            combine(diffWaysToCompute(expression[:i]),
                    diffWaysToCompute(expression[i + 1:]),
                    &res,
                    func(a, b int) int { return a + b })
        case '-':
            combine(diffWaysToCompute(expression[:i]),
                    diffWaysToCompute(expression[i + 1:]),
                    &res,
                    func(a, b int) int { return a - b })                
        case '*':
            combine(diffWaysToCompute(expression[:i]),
                    diffWaysToCompute(expression[i + 1:]),
                    &res,
                    func(a, b int) int { return a * b })
        default:
            num = 10 * num + int(expression[i] - '0')
        }
    }
    if len(res) == 0 {
        res = append(res, num)
    }
    return res
}

func combine(l, r []int, res *[]int, f func(int, int) int) {
    for _, a := range l {
        for _, b := range r {
            *res = append(*res, f(a, b))
        }
    }
}