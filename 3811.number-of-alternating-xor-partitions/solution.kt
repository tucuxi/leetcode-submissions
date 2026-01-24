private const val MOD = 1_000_000_007
private const val MAX_VAL = 200_000 

class Solution {
    fun alternatingXOR(nums: IntArray, target1: Int, target2: Int): Int {
        val dp1 = IntArray(MAX_VAL)
        val dp2 = IntArray(MAX_VAL).also { it[0] = 1 }
        var currentXor = 0
        var waysToEndHere = 0

        nums.forEach { num ->
            currentXor = currentXor xor num
            val waysToFormTarget1 = dp2[currentXor xor target1]
            val waysToFormTarget2 = dp1[currentXor xor target2]
            dp1[currentXor] = (dp1[currentXor] + waysToFormTarget1) % MOD
            dp2[currentXor] = (dp2[currentXor] + waysToFormTarget2) % MOD
            waysToEndHere = (waysToFormTarget1 + waysToFormTarget2) % MOD
        }

        return waysToEndHere
    }
}