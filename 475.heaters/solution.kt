class Solution {
    fun findRadius(houses: IntArray, heaters: IntArray): Int {
        var res = 0
        
        heaters.sort()

        for (house in houses) {
            val i = heaters.binarySearch(house).let {
                if (it < 0) -(it + 1) else it
            }
            val distances = sequence {
                if (i > 0) {
                    yield(house - heaters[i - 1])
                }
                if (i < heaters.size) {
                    yield(heaters[i] - house)
                }
            }
            res = maxOf(res, distances.min())
        }
        
        return res
    }
}