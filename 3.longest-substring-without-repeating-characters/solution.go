func lengthOfLongestSubstring(s string) int {
    lastOccurrence := map[byte]int{}
    run, longestRun := 0, 0
    for i := range s {
        if j, ok := lastOccurrence[s[i]]; ok {
            if i - j <= run {
                run = i - j - 1
            }
        }
        lastOccurrence[s[i]] = i
        run++
        if run > longestRun {
            longestRun = run
        }
    }
    return longestRun
}