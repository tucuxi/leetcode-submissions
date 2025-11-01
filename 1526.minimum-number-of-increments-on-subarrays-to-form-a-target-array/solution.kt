class Solution {
    fun minNumberOperations(target: IntArray): Int {
        var res = target[0]
        for (i in 1 .. target.lastIndex) {
            res += maxOf(target[i] - target[i - 1], 0)
        }
        return res
    }
}