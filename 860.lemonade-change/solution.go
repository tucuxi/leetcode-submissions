func lemonadeChange(bills []int) bool {
    fives, tens := 0, 0
    for _, bill := range bills {
        switch bill {
            case 5:
                fives++
            case 10:
                fives--
                tens++
            case 20:
                if tens > 0 {
                    fives--
                    tens--
                } else {
                    fives -= 3
                }
        }
        if fives < 0 {
            return false
        }
    }
    return true
}