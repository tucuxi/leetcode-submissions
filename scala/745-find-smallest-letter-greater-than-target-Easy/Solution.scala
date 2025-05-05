object Solution {
    private val nchars = 'z' - 'a' + 1

    // missed that array is sorted
    
    def nextGreatestLetter(letters: Array[Char], target: Char): Char = {
        def dist(c: Char): Int =
            if (c > target) c - target else c - target + nchars
    
        var minDist = nchars
        var minChar = target
        for (c <- letters)
            if (dist(c) < minDist) {
                minDist = dist(c)
                minChar = c
            }
        minChar
    }
}