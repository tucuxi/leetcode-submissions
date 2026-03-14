class Solution {
    fun getHappyString(n: Int, k: Int): String {
        return generate("", n).drop(k - 1).firstOrNull().orEmpty()  
    }

    fun generate(pre: String, n: Int): Sequence<String> {
        return sequence {
            if (pre.length == n) {
                yield(pre)
            } else {
                "abc".forEach { ch ->
                    if (pre.lastOrNull() != ch) {
                        yieldAll(generate(pre + ch, n))
                    }
                }
            }
        }
    }
}