object Solution {
    def numJewelsInStones(J: String, S: String): Int = {
        var res = 0
        for (s <- S if J.indexOf(s) >= 0) res += 1
        res
    }
}