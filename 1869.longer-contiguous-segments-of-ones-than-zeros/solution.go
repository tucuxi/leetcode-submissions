func checkZeroOnes(s string) bool {
    max := [2]int{}
    cur := 1
    ind := int(s[0] - '0')
    max[ind] = cur
    for i := 1; i < len(s); i++ {
        if s[i] == s[i - cur] {
            cur++
        } else {
            cur = 1
            ind = int(s[i] - '0')
        }
        if cur > max[ind] {
            max[ind] = cur 
        }
    }
    return max[1] > max[0]
}