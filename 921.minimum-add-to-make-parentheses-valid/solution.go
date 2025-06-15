func minAddToMakeValid(s string) int {
    insert := 0
    balance := 0
    for _, ch := range s {
        if ch == '(' {
            balance++
        } else if balance == 0 {
            insert++
        } else {
            balance--
        }
    }
    return insert + balance 
}