class Solution {
    fun mySqrt(x: Int): Int {
        if (x <= 1) return x
        var x0 = x / 2
        while (true) {
	        val x1 = (x0 + x / x0) / 2
		    if (x1 >= x0) return x0
			x0 = x1
		}
    }
}