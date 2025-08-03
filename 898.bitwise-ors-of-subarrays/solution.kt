class Solution {
    fun subarrayBitwiseORs(arr: IntArray): Int {
        val res = mutableSetOf<Int>()
        var acc = setOf<Int>()
        arr.forEach { n ->
            acc = setOf(n) + acc.map { it or n }
            res += acc
        }
        return res.size
    }
}