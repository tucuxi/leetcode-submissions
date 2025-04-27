func robotSim(commands []int, obstacles [][]int) int {
    h := make(map[[2]int]bool)
    for _, o := range obstacles {
        h[[2]int{o[0], o[1]}] = true
    }
    x, y := 0, 0
    dx, dy := 0, 1
    maxDist := 0
    for _, c := range commands {
        switch c {
        case -2:
            dx, dy = -dy, dx
        case -1:
            dx, dy = dy, -dx
        default:
            for ; c > 0; c-- {
                x2, y2 := x + dx, y + dy
                if h[[2]int{x2, y2}] {
                    break
                }
                x, y = x2, y2
                maxDist = max(maxDist, x * x + y * y)
            }            
        }         
    }
    return maxDist
}
