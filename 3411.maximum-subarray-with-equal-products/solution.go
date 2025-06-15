func maxLength(nums []int) int {
    res := 1
    for i, m := range nums {
        p := m
        g := m
        l := m
        for j, n := range nums[i + 1:] {
            p *= n
            g = gcd(g, n)
            l = lcm(l, n)
            if p != g * l {
                break
            }
            res = max(res, j + 2)
        }
    }
    return res
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func lcm(a, b int) int {
    return a / gcd(a, b) * b
}
