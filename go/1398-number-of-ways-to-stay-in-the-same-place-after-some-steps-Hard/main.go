func numWays(steps int, arrLen int) int {
    const mod = 1_000_000_007

    maxDistance := func() int {
        if arrLen > steps {
            return steps
        }
        return arrLen
    }()
    
    cur := make([]int, maxDistance)
    prev := make([]int, maxDistance)
    cur[0] = 1
 
    for remainingSteps := 1; remainingSteps <= steps; remainingSteps++ {
        prev, cur = cur, prev
        for pos := 0; pos < maxDistance; pos++ {
            n := prev[pos]
            if pos > 0 {
                n += prev[pos - 1]
            }
            if pos + 1 < maxDistance {
                n += prev[pos + 1]
            }
            cur[pos] = n % mod
        }
    }
    
    return cur[0]
}