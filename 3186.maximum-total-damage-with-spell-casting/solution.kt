class Solution {
    fun maximumTotalDamage(power: IntArray): Long {
        val counts = power.asIterable().groupingBy { it }.eachCount()
        val keys = counts.keys.sorted()

        val dp = LongArray(keys.size)
        var previousMaxDamage = 0L
        var maxDamage = 0L
        var i = 0

        for (j in keys.indices) {
            val spell = keys[j]
            while (i < j && keys[i] < spell - 2) {
                previousMaxDamage = maxOf(previousMaxDamage, dp[i])
                i++
            }
            dp[j] = previousMaxDamage + spell.toLong() * counts.getValue(spell)
            maxDamage = maxOf(dp[j], maxDamage)
        }

        return maxDamage
    }
}