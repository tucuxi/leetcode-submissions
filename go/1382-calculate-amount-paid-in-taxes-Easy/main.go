func calculateTax(brackets [][]int, income int) float64 {
    tax := 0.0
    taxed := 0
    lower := 0
    for i := 0; taxed < income; i++ {
        var within int
        if income > brackets[i][0] {
            within = brackets[i][0] - lower
        } else {
            within = income - lower
        }
        lower = brackets[i][0]
        tax += float64(within) * float64(brackets[i][1]) / 100.0
        taxed += within
    }
    return tax
}