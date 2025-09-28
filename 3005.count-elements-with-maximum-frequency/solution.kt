class Solution {
    fun maxFrequencyElements(nums: IntArray): Int {
        val histogram = IntArray(101)
        var maxFrequency = 0
        var result = 0
        nums.forEach { n ->
            histogram[n]++
            if (histogram[n] == maxFrequency) {
                result += maxFrequency
            } else if (histogram[n] > maxFrequency) {
                maxFrequency = histogram[n]
                result = maxFrequency
            }
        }
        return result
    }
}