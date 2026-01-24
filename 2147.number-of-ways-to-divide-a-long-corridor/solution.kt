const val MOD = 1_000_000_007

class Solution {
    fun numberOfWays(corridor: String): Int {
        var res = 1L
        var s = 0
        var p = 0
        
        corridor.forEach { c ->
            when (c) {

                'S' -> {
                    if (s < 2) {
                        s++
                    } else {
                        res = res * (p + 1) % MOD
                        s = 1
                        p = 0
                    }
                }

                'P' -> {
                    if (s == 2) {
                        p++
                    }
                }
            }
        }
        return if (s < 2) 0 else res.toInt() 
    }
}