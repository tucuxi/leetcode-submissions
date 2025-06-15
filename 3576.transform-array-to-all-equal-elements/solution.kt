class Solution {
    fun canMakeEqual(nums: IntArray, k: Int): Boolean {
        
        fun transform(target: Int): Boolean {
            val (count, lastTarget) = nums
                .dropLast(1)
                .fold(0 to target) { (c, t), num ->
                    if (num == t) {
                        Pair(c, target)
                    } else {
                        Pair(c + 1, -target)
                    }
                }
            return count <= k && lastTarget == nums.last() 
        }

        return transform(1) || transform(-1)
    }
}