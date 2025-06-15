object Solution {
    def bitwiseComplement(N: Int): Int = {
        def mask(n: Int): Int = {
            var m = 1
            var p = n
            while ((p >>> 1) > 0) {
                m = (m << 1) | 1
                p >>>= 1
            }
            m
        }
        (N ^ -1) & mask(N)
    }
}