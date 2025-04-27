class Solution {
    fun trimMean(arr: IntArray): Double {
        val k = arr.size / 20
        val res = arr.sorted().subList(k, arr.size - k)
        return res.average()
    }
}