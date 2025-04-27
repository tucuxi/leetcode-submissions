func distinctIntegers(n int) int {
    set := map[int]bool{n: true}
    for {
        size := len(set)
        for i := 1; i <= n; i++ { 
            for x := range set {
                if x % i == 1 {
                    set[i] = true
                }             
            }
        }
        if size == len(set) {
            break
        }
    }
    return len(set)
}