const keys = 8

func minimumPushes(word string) int {
    var h [26]int
    
    for _, ch := range word {
        h[ch - 'a']++
    }
    
    var f []int
    
    for _, n := range h {
        if n > 0 {
            f = append(f, n)
        }
    }
    
    slices.Sort(f)
    
    res := 0
    k := 0
    w := 1
    
    for i := len(f) - 1; i >= 0; i-- {
        if k == keys {
            k = 0
            w++
        }
        k++
        res += w * f[i]
    }
    
    return res
}