/** 
 * The API guess is defined in the parent class.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * fun guess(num:Int):Int {}
 */

class Solution:GuessGame() {
    override fun guessNumber(n:Int):Int {
        var low = 1
        var high = n
        while (low < high) {
            val mid = low + (high - low) / 2
            when(guess(mid)) {
                -1 -> high = mid - 1
                0 -> return mid
                1 -> low = mid + 1
            }
        }
        return low
    }
}