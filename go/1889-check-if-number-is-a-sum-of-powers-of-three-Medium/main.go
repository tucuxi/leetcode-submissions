const base = 3

func checkPowersOfThree(n int) bool {
    for ; n > 0; n /= base {
        if n % base > 1 {
            return false
        }
    }
    return true
}