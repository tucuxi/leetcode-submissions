func canConstruct(ransomNote string, magazine string) bool {
    ransomHist := count(ransomNote)
    magazineHist := count(magazine)
    for i := range ransomHist {
        if ransomHist[i] > magazineHist[i] {
            return false
        }
    }
    return true
}

func count(s string) [26]int {
    res := [26]int{}
    for i := range s {
        res[s[i] - 'a']++
    }
    return res
}