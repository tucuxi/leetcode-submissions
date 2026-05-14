func sumOfPrimesInRange(n int) int {
    r := reverse(n)
    p := min(n, r)
    q := max(n, r)
    
    sieve := make([]bool, q+1)
    sieve[1] = true

    for i := 4; i <= q; i += 2 {
        sieve[i] = true
    }
    for i := 3; i <= q; i++ {
        for j := 2*i; j <= q; j += i {
            sieve[j] = true
        }
    }

    sum := 0

    for i := p; i <= q; i++ {
        if !sieve[i] {
            sum += i
        }
    }

    return sum
}

func reverse(n int) int {
    r := 0
    for ; n > 0; n /= 10 {
        r = 10*r + n%10
    }
    return r
}