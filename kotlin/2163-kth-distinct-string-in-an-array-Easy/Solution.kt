class Solution {
    fun kthDistinct(arr: Array<String>, k: Int): String {
        val count = arr.groupingBy { it }.eachCount()
        return arr.filter { count[it]!! == 1 }.elementAtOrElse(k - 1) { "" }
    }
}