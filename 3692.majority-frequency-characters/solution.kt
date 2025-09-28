class Solution {
    fun majorityFrequencyGroup(s: String): String {
        val charToFrequency = s.groupingBy { it }.eachCount().asIterable()
        val frequencyToCount = charToFrequency.groupingBy { it.value }.eachCount()
        val majorityFrequency = frequencyToCount.maxWith(compareBy<Map.Entry<Int, Int>>({ it.value }, { it.key })).key
        return charToFrequency
            .filter { it.value == majorityFrequency }
            .joinToString("") { it.key.toString() }
    }
}