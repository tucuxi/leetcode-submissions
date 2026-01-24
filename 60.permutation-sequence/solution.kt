class Solution {
    fun getPermutation(n: Int, k: Int): String {

        val chars = CharArray(n) { '1' + it }
        val used = BooleanArray(n)
        val b = StringBuilder()

        fun generate(): Sequence<String> = sequence {
            if (b.length == n) {
                yield(b.toString())
            } else {
                for (i in 0 until n) {
                    if (!used[i]) {
                        used[i] = true
                        b.append(chars[i])
                        yieldAll(generate())
                        b.setLength(b.length - 1)
                        used[i] = false
                    }
                }
            }
        }

        return generate().elementAt(k - 1)
    }
}