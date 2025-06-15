class Solution {
    fun readBinaryWatch(turnedOn: Int): List<String> {
        val res = mutableListOf<String>()
        for (h in 0..11) {
            for (m in 0..59) {
                if (Integer.bitCount(h) + Integer.bitCount(m) == turnedOn) {
                    res.add("%d:%02d".format(h, m))
                } 
            }
        }
        return res
    }
}