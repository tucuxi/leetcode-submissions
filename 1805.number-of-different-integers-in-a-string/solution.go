func numDifferentIntegers(word string) int {
    var b strings.Builder
    for i := range word {
        if word[i] >= 'a' && word[i] <= 'z' {
            b.WriteByte(' ')
        } else {
            b.WriteByte(word[i])
        }
    }
    modifiedWord := b.String()
    nums := map[string]bool{}
    for _, number := range strings.Fields(modifiedWord) {
        for ; len(number) > 1 && number[0] == '0'; number = number[1:] {
        }
        nums[number] = true
    }
    return len(nums)
}