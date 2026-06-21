class Solution {
    fun internalAngles(sides: IntArray): DoubleArray {
        val (a, b, c) = sides.sorted()
        
        if (a + b <= c) {
            return doubleArrayOf()
        }

        val alpha = acos((b*b + c*c - a*a).toDouble() / (2*b*c).toDouble())

        val beta = acos((a*a + c*c - b*b).toDouble() / (2*a*c).toDouble())

        return doubleArrayOf(alpha/PI * 180, beta/PI * 180, (1 - alpha/PI - beta/PI) * 180)
    }
}