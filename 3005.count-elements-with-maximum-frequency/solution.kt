class Solution {
    fun maxFrequencyElements(nums: IntArray): Int {
        val h = nums.asIterable()
            .groupingBy { it }
            .eachCount()
            .values
            .sortedByDescending { it }
        val maxFrequency = h.first()
        val m = h.takeWhile { it == maxFrequency }
        return m.size * maxFrequency        
    }
}