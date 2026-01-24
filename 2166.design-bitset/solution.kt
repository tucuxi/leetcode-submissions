class Bitset(size: Int) {

    private val bits = BooleanArray(size)
    private var flipped = false 
    private var count0 = size

    fun fix(idx: Int) {
        if (bits[idx] == flipped) {
            bits[idx] = !flipped
            count0--
        } 
    }

    fun unfix(idx: Int) {
        if (bits[idx] != flipped) {
            bits[idx] = flipped
            count0++
        }
    }

    fun flip() {
        flipped = !flipped
        count0 = bits.size - count0
    }

    fun all(): Boolean {
        return count0 == 0
    }

    fun one(): Boolean {
        return count0 < bits.size
    }

    fun count(): Int {
        return bits.size - count0
    }

    override fun toString(): String {
        return bits.joinToString("") {
            if (it != flipped) "1" else "0"
        }
        
    }

}

/**
 * Your Bitset object will be instantiated and called as such:
 * var obj = Bitset(size)
 * obj.fix(idx)
 * obj.unfix(idx)
 * obj.flip()
 * var param_4 = obj.all()
 * var param_5 = obj.one()
 * var param_6 = obj.count()
 * var param_7 = obj.toString()
 */