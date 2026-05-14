var change = map[int]bool{2: true, 5: true, 6: true, 9: true}

func rotatedDigits(n int) int {
    res := 0

    outer:
    for i := 1; i <= n; i++ {
        c := false
        for k := i; k > 0; k /= 10 {
            d := k%10
            if d == 3 || d == 4 || d == 7 {
                continue outer
            }
            if d == 2 || d == 5 || d == 6 || d == 9 {
                c = true
            }
        }
        if c {
            res++
        }
    }
    
    return res
}
