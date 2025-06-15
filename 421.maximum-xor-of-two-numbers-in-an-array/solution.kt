class Solution {
    fun findMaximumXOR(nums: IntArray): Int {
        val bits = generateSequence(1 shl 31) {
            if (it != 0) it ushr 1 else null
        }
        val (res, _) = bits
            .fold(0 to 0) { (maxXor, mask), bits ->
                val possibleMaxXor = maxXor or bits
                val newMask = mask or bits
                val maskedNums = nums.map { it and newMask }.toSet()
                if (maskedNums.any { m -> m xor possibleMaxXor in maskedNums }) {
                    possibleMaxXor to newMask
                } else {
                    maxXor to newMask
                }
            }
        return res
    }
}