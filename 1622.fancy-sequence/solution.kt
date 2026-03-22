const val MOD = 1_000_000_007

class Fancy() {
    private val values = mutableListOf<Int>()
    private var factor = 1L
    private var term = 0L
        
    fun append(`val`: Int) {
        val v = (`val`.toLong() - term + MOD) % MOD * inv(factor.toInt()) % MOD
        values += v.toInt()
    }

    fun addAll(inc: Int) {
        term = (term + inc) % MOD
    }

    fun multAll(m: Int) {
        term = term * m % MOD
        factor = factor * m % MOD
    }

    fun getIndex(idx: Int): Int {
        return if (idx > values.lastIndex) {
            return -1
        } else {
            ((factor * values[idx] % MOD + term) % MOD).toInt()
        }
    }

    private fun quickmul(x: Int, y: Int): Int {
        var res = 1L
        var cur = x.toLong()
        var y = y

        while (y != 0) {
            if (y and 1 != 0) {
                res = res * cur % MOD
            }
            cur = cur * cur % MOD
            y = y shr 1
        }
        return res.toInt()
    }

    private fun inv(x: Int): Int {
        return quickmul(x, MOD - 2)
    }
}

/**
 * Your Fancy object will be instantiated and called as such:
 * var obj = Fancy()
 * obj.append(`val`)
 * obj.addAll(inc)
 * obj.multAll(m)
 * var param_4 = obj.getIndex(idx)
 */