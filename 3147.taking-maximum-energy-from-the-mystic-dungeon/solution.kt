class Solution {
    fun maximumEnergy(energy: IntArray, k: Int): Int {
        val runningTotal = IntArray(k)
        var maxEnergy = Int.MIN_VALUE
        for (i in energy.lastIndex downTo 0) {
            val j = i % k
            runningTotal[j] += energy[i]
            maxEnergy = maxOf(runningTotal[j], maxEnergy)
        }
        return maxEnergy
    }
}