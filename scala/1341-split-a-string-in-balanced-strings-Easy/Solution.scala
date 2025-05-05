object Solution {
    def balancedStringSplit(s: String): Int = {
        var bal = 0
        var cnt = 0
        for (c <- s) {
            c match {
                case 'L' => bal += 1
                case 'R' => bal -= 1
            }
            if (bal == 0) cnt += 1
        }
        cnt
    }
}