class Solution {
    fun canPlaceFlowers(flowerbed: IntArray, n: Int): Boolean {
        if (n == 0) return true
        var m = n
        var empty = 1
        flowerbed.forEach { plot ->
            if (plot == 0) {
                if (++empty == 3) {
                    if (--m == 0) {
                        return true
                    }
                    empty = 1
                }
            } else {
                empty = 0
            }
        }
        return m == 1 && empty == 2
    }
}