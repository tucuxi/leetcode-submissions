object Solution {
    def hammingDistance(x: Int, y: Int): Int = {
        var p = 1
        var res = 0
        for (i <- 0 until 32) {
            if ((x & p) != (y & p)) res += 1
            p <<= 1
        }
        res
    }
}