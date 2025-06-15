/* The guess API is defined in the parent class GuessGame.
   @param num, your guess
   @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
      int guess(int num); */

public class Solution extends GuessGame {
    public int guessNumber(int n) {
        int lo = 1;
        int hi = n;
        while (lo <= hi) {
            int k = lo + (hi - lo) / 2;
            switch (guess(k)) {
                case -1:
                    hi = k -1;
                    break;
                case 0:
                    return k;
                case 1:
                    lo = k + 1;
                    break;
            }
        }
        return -1;
    }
}