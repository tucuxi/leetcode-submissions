func maxActiveSectionsAfterTrade(s string) int {
    ones := 0

    for _, c := range s {
        if c == '1' {
            ones++
        }
    }

    var zeroBlocks []int

    for i := 0; i < len(s); {
        j := i
        for i < len(s) && s[i] == s[j] {
            i++
        }
        if s[j] == '0' {
            zeroBlocks = append(zeroBlocks, i-j)
        }
    }
    
    gain := 0
    
    for i := range len(zeroBlocks) - 1 {
        gain = max(gain, zeroBlocks[i] + zeroBlocks[i+1])
    }

    return ones + gain
}