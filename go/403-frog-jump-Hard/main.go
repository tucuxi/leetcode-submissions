func canCross(stones []int) bool {
    advance := func(currentIndex, distance int) (int, bool) {
        if distance > 0 {
            targetPosition := stones[currentIndex] + distance
            for index := currentIndex + 1; index < len(stones); index++ {
                if stones[index] > targetPosition {
                    break
                }
                if stones[index] == targetPosition {
                    return index, true
                }
            }
        }
        return -1, false
    }
    
    memo := make(map[[2]int]bool)
    
    var dp func(int, int) bool
    dp = func(currentIndex, previousJump int) bool {
        if currentIndex == len(stones) - 1 {
            return true
        }
        if res, exists := memo[[2]int{currentIndex, previousJump}]; exists {
            return res
        }
        for jump := previousJump - 1; jump <= previousJump + 1; jump++ {
            nextIndex, exists := advance(currentIndex, jump)
            if exists && dp(nextIndex, jump) {
                memo[[2]int{currentIndex, previousJump}] = true
                return true
            }
        }
        memo[[2]int{currentIndex, previousJump}] = false
        return false
    }
    
    return dp(0, 0)
}