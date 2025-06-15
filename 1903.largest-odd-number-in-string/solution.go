func largestOddNumber(num string) string {
    i := len(num) - 1
    for ; i >= 0 && strings.IndexByte("13579", num[i]) == -1; i-- {}
    return num[:i+1]
}