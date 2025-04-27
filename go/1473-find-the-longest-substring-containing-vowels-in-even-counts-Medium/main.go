func findTheLongestSubstring(s string) int {
    characterMap := make(map[byte]int)
    characterMap['a'] = 1 << 0
    characterMap['e'] = 1 << 1
    characterMap['i'] = 1 << 2
    characterMap['o'] = 1 << 3
    characterMap['u'] = 1 << 4
    mp := [32]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
    prefixXor := 0
    longestSubstring := 0
    for i := range s {
        prefixXor ^= characterMap[s[i]]
        if prefixXor != 0 && mp[prefixXor] == -1 {
            mp[prefixXor] = i
        }
        longestSubstring = max(longestSubstring, i - mp[prefixXor])
    }
    return longestSubstring
}