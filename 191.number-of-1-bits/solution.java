public class Solution {
    // you need to treat n as an unsigned value
    public int hammingWeight(int n) {
        int res = 0;
        while (n != 0) {
            if ((n & 1) != 0) res++;
            n >>>= 1;
        }
        return res;
    }
}