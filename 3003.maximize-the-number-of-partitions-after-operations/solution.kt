class Solution {
    fun maxPartitionsAfterOperations(s: String, k: Int): Int {
        val sf = Array(s.length) { IntArray(2) }
        var msk = 0
        var part = 0
        var res = 0
        
        for (i in s.lastIndex downTo 1) {
            val bit = 1 shl (s[i]-'a')
            msk = msk or bit
            if (msk.countOneBits() > k) {
                part++
                msk = bit
            }
            sf[i - 1][0] = part
            sf[i - 1][1] = msk
        }
        
        msk = 0
        part = 0
        
        for (i in 0 until s.lastIndex) {
            val cntall = (msk or sf[i][1]).countOneBits()
            var c = part + sf[i][0]
            if (msk.countOneBits() == k && sf[i][1].countOneBits() == k && cntall < 26) {
                c += 2
            }
            else if (minOf(cntall + 1, 26) > k) {
                c++
            }
            res = maxOf(res, c)
            val bit = 1 shl (s[i]-'a')
            msk = msk or bit
            if (msk.countOneBits() > k) {
                part++
                msk = bit
            }
        }
        return res + 1
    }
}