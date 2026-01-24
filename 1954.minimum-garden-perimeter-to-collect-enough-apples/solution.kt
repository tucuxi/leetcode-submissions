class Solution {
    fun minimumPerimeter(neededApples: Long): Long {
        var l = 1L
        var r = 100000L

        while (l < r) {
            val m = l + (r - l) /2
            val apples = 4 * m * m * m + 6 * m * m + 2 * m
            if (apples < neededApples) {
                l = m + 1
            } else {
                r = m
            }
        }

        return 8 * l
     }
}