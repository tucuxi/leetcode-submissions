func compress(chars []byte) int {
    i, j, k := 0, 0, 0
    finish := func() {
        chars[k] = chars[i]
        k++
        if rl := j - i; rl > 1 {
            for _, c := range strconv.Itoa(rl) {
                chars[k] = byte(c)
                k++
            }
        }
        i = j
    }
    for ; j < len(chars); j++ {
        if chars[j] != chars[i] {
            finish()
        }
   } 
   finish()
   return k
}