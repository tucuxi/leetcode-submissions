object Solution {
    def isOneBitCharacter(bits: Array[Int]): Boolean = {
        var i = 0
        var one = true
        while (i < bits.length) {
            one = bits(i) == 0
            if (one) i += 1 else i += 2 
        }
        one
    }
}