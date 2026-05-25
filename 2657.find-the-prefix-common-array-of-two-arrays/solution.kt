class Solution {
    fun findThePrefixCommonArray(A: IntArray, B: IntArray): IntArray {
        val n = A.size
        val va = BooleanArray(n + 1)
        val vb = BooleanArray(n + 1)
        var c = 0

        return IntArray(n) { i ->
            A[i].let { a ->
                if (!va[a]) {
                    va[a] = true
                    if (vb[a]) {
                        c++
                    }
                }
            }
            B[i].let { b ->
                if (!vb[b]) {
                    vb[b] = true
                    if (va[b]) {
                        c++
                    }
                }
            }
            c
        }
    }
}
