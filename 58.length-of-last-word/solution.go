func lengthOfLastWord(s string) int {
    j := len(s)-1
    
    for s[j] == ' ' {
        j--
    }
    
    i := j

    for i > 0 && s[i-1] != ' ' {
        i--
    }

    return j-i+1
}