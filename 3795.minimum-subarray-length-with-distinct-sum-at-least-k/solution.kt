class Solution {
    fun minLength(nums: IntArray, k: Int): Int {
        val h = IntArray(100001)
        var p = 0
        var q = 0
        var sum = 0
        var res = Int.MAX_VALUE

        while (p < nums.size) {
            when {
                sum >= k -> {
                    res = minOf(res, q - p)
                    nums[p].let { np ->
                        if (--h[np] == 0) {
                            sum -= np
                        }
                    }
                    p++
                }
                
                q < nums.size -> {
                    nums[q].let { nq ->
                        if (++h[nq] == 1) {
                            sum += nq
                        }
                    }
                    q++
                }
                
                else -> break
            }
        }
        return if (res == Int.MAX_VALUE) -1 else res
    }
}