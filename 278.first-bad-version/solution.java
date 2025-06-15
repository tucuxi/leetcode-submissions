/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */

public class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        return bisect(1, n);
    }
    private int bisect(int lo, int hi) {
        if (lo >= hi) return hi;
        int mi = lo + (hi - lo) / 2;
        return isBadVersion(mi) ? bisect(lo, mi) : bisect(mi + 1, hi);
    }
}