class Solution {
    fun sortByReflection(nums: IntArray): IntArray {
        return nums.sortedWith(
            compareBy(
                { reflect(it) },
                { it }
            )
        ).toIntArray()        
    }

    fun reflect(num: Int): Int {
        var r = 0
        var n = num
        while (n > 0) {
            r = 2*r + n%2
            n /= 2
        }
        return r
    }
}