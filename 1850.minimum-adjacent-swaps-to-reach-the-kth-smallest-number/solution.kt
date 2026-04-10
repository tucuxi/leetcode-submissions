class Solution {
    fun getMinSwaps(num: String, k: Int): Int {
        val target = num.toCharArray()

        repeat(k) {
            var i = target.lastIndex - 1
            while (i >= 0 && target[i] >= target[i + 1]) {
                i--
            }
            if (i >= 0) {
                var j = target.lastIndex
                while (target[j] <= target[i]) {
                    j--
                }
                target[i] = target[j].also { target[j] = target[i] }
                target.reverse(i + 1, target.size)
            }
        }

        val source = num.toCharArray()
        var res = 0

        for (i in source.indices) {
            if (source[i] != target[i]) {
                var j = i + 1
                while (source[j] != target[i]) {
                    j++
                }
                while (j > i) {
                    source[j - 1] = source[j].also { source[j] = source[j - 1] }
                    res++
                    j--
                }
            }
        }

        return res
    }
}