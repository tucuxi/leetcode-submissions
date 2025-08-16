class Solution {
    fun largestGoodInteger(num: String): String {
        var run = 1
        var res: Char? = null
        for (i in 1 until num.length) {
            if (num[i] == num[i - 1]) {
                run++
                if (run == 3 && (res == null || res < num[i])) {
                    res = num[i]
                }
            } else {
                run = 1
            }
        }
        return res?.let { c -> "$c$c$c" } ?: ""
    }
}