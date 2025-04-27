func maxPoints(points [][]int) int {
    if len(points) <= 2 {
        return len(points)
    }
    max := 0
    for i, p := range points {
        h := map[[2]int]int{}
        for j, q := range points {
            if j == i {
                continue
            }
            dx := p[0] - q[0]
            dy := p[1] - q[1]
            if div := gcd(dx, dy); div != 0 {
                dx /= div
                dy /= div
            }
            xy := [2]int{dx, dy}
            h[xy]++
            if h[xy] > max {
                max = h[xy]
            }
        }
    }
    return max + 1
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}