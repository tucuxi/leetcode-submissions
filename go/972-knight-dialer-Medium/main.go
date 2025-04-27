func knightDialer(n int) int {
    const mod = 1e9 + 7
    b1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
    b2 := make([]int, 10)
    
    for i := 1; i < n; i++ {
        copy(b2, b1)
        b1[0] = (b2[4] + b2[6]) % mod
        b1[1] = (b2[6] + b2[8]) % mod
        b1[2] = (b2[7] + b2[9]) % mod
        b1[3] = (b2[4] + b2[8]) % mod
        b1[4] = (b2[0] + b2[3] + b2[9]) % mod
        b1[5] = 0
        b1[6] = (b2[0] + b2[1] + b2[7]) % mod
        b1[7] = (b2[2] + b2[6]) % mod
        b1[8] = (b2[1] + b2[3]) % mod
        b1[9] = (b2[2] + b2[4]) % mod
    }
    
    res := 0
    for _, c := range b1 {
        res = (res + c) % mod 
    }
    return res
}