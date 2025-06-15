func getWordsInLongestSubsequence(words []string, groups []int) []string {
    n := len(words)
    dp := make([]int, n)
    pred := make([]int, n)
    k := 0
    
    for i := range n {
        dp[i] = 1
        pred[i] = -1
    }

    for i := 1; i < n; i++ {
        for j := range i {
            if groups[i] != groups[j] &&
                    len(words[i]) == len(words[j]) &&
                    hamming(words[i], words[j]) == 1 &&
                    dp[i] < dp[j] + 1 {
                dp[i] = dp[j] + 1
                pred[i] = j
                if dp[i] > dp[k] {
                    k = i
                }
            }
        }
    }

    var res []string
    for k >= 0 {
        res = append(res, words[k])
        k = pred[k]
    }
    l := len(res) - 1
    for i := range l / 2 + 1 {
        res[i], res[l - i] = res[l - i], res[i]
    } 
    return res
}

func hamming(s, t string) int {
    res := 0
    for i := range s {
        if s[i] != t[i] {
            res++
        }
    }
    return res
}