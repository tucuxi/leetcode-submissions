func simplifiedFractions(n int) []string {
    var res []string

    for nom := 1; nom < n; nom++ {
        for denom := nom + 1; denom <= n; denom++ {
            if gcd(nom, denom) == 1 {
                res = append(res, fmt.Sprintf("%d/%d", nom, denom))
            }
        }
    }
    
    return res
}

func gcd(a, b int) int {
    for b > 0 {
        a, b = b, a % b
    }
    return a
}