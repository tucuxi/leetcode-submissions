func canMakeSubsequence(str1 string, str2 string) bool {
    j := 0
    for i := 0; i < len(str1) && j < len(str2); i++ {
        if str1[i] == str2[j] || (str2[j] - str1[i] + 26) % 26 == 1 {
            j++
        }
    }
    return j == len(str2)
}