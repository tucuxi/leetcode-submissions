func canThreePartsEqualSum(arr []int) bool {
    total := 0
    for _, v := range arr {
        total += v
    }
    if total % 3 != 0 {
        return false
    }
    parts := 0
    sum := 0
    for _, v := range arr {
        sum += v
        if sum == total / 3 {
            sum = 0
            parts++
            if parts == 3 {
                return true
            }
        }
    }
    return false
}