object Solution {
    def judgeCircle(moves: String): Boolean = {
        moves.count(_ == 'D') == moves.count(_ == 'U') &&
            moves.count(_ == 'L') == moves.count(_ == 'R')
        
        /*
        var x, y = 0
        for (m <- moves) m match {
            case 'D' => y -= 1
            case 'U' => y += 1
            case 'L' => x -= 1
            case 'R' => x += 1
        }
        x == 0 && y == 0
        */
    }
}