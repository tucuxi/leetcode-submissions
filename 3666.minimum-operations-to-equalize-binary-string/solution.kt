class Solution {
    fun minOperations(s: String, k: Int): Int {
        val zeros = s.count { it == '0' }

        if (zeros == 0) {
            return 0
        }

        val n = s.length

        if (n == k) {
            return if (zeros == n) 1 else -1
        }

        val ones = n - zeros
        val b = n - k

        return sequence {
            if (zeros % 2 == k % 2) {
                var m = maxOf((zeros + k - 1) / k, (ones + b - 1) / b)
                m += m.inv() and 1
                yield(m)
            }
            if (zeros % 2 == 0) {
                var m = maxOf((zeros + k - 1) / k, (zeros + b - 1) / b)
                m += m and 1
                yield(m)
            }
        }
        .minOrNull() ?: -1
    }
}