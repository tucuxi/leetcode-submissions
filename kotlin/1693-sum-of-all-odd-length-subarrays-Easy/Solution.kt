class Solution {
    fun sumOddLengthSubarrays(arr: IntArray): Int {
        var res = 0
        for (i in 1..arr.size step 2)
            for (j in 0..arr.size - i)
                for (k in j until j + i)
                    res += arr[k]
        return res
    }
}