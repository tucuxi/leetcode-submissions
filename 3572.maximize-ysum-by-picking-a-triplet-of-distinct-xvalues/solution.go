func maxSumDistinctTriplet(x []int, y []int) int {
    g := make(map[int]int)

    for i, xi := range x {
        g[xi] = max(g[xi], y[i]) 
    }

    if len(g) < 3 {
        return -1
    }

    y1, y2, y3 := 0, 0, 0

    for _, yi := range g {
        if yi > y1 {
            y1, y2, y3 = yi, y1, y2
        } else if yi > y2 {
            y2, y3 = yi, y2
        } else if yi > y3 {
            y3 = yi
        }
    }

    return y1+y2+y3
}