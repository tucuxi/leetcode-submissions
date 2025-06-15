func maximumSwap(num int) int {
    digits := []byte(strconv.Itoa(num))
    index := make([]int, len(digits))
    for i := range index {
        index[i] = i
    }
    slices.SortFunc(index, func(a, b int) int {
        d := int(digits[b]) - int(digits[a])
        if d == 0 {
            return a - b
        }
        return d
    })
    for i := range index {
        if i != index[i] {
            j := i 
            for j + 1 < len(index) && digits[index[j + 1]] == digits[index[i]] {
                j++
            }
            digits[i], digits[index[j]] = digits[index[j]], digits[i]
            res, _ := strconv.Atoi(string(digits))
            return res
        }
    }
    return num
}