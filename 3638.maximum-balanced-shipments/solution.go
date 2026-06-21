func maxBalancedShipments(weight []int) int {
    n := len(weight)
    res := 0

    for i := 0; i < n; {
        m := weight[i]
        
        for i++; i < n && weight[i] >= m; i++ {
            m = weight[i]
        }
        if i < n {
            res++
        }
        i++
    }

    return res
}