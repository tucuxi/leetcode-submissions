func largestVariance(s string) int {
    freq := make(map[rune]int)
    for _, ch := range s {
        freq[ch]++
    }
        
    results := make(chan int, 8)
    
    variance := func(ch1, ch2 rune) {
        res := 0
        f1, f2, r2 := 0, 0, freq[ch2]
        for _, ch := range s {
            if ch == ch1 {
                f1++
            } else if ch == ch2 {
                f2++
                r2--
            }
            if f2 > 0 {
                if v := f1 - f2; v > res {
                    res = v
                }
            }
            if f1 < f2 && r2 > 0 {
                f1, f2 = 0, 0
            }
        }
        results <- res        
    }

    pairs := 0
    
    for ch1 := 'a'; ch1 <= 'z'; ch1++ {
        for ch2 := 'a'; ch2 <= 'z'; ch2++ {
            if ch1 != ch2 && freq[ch1] > 0 && freq[ch2] > 0 {
                pairs++
                go variance(ch1, ch2)
            }
        }
    }

    maxVariance := 0
    
    for ; pairs > 0; pairs-- {
        if result := <-results; result > maxVariance {
            maxVariance = result
        }
    }
    
    return maxVariance
}
