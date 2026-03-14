class Solution {
    fun gcdSum(nums: IntArray): Long {
        var mx = nums[0]

        val prefixGcd = IntArray(nums.size) {
            mx = maxOf(mx, nums[it])
            gcd(nums[it], mx)
        }

        prefixGcd.sort()

        val sum = (0 until nums.size / 2).sumOf {
            gcd(prefixGcd[it], prefixGcd[prefixGcd.lastIndex - it]).toLong()
        }

        return sum
    }

    private tailrec fun gcd(a: Int, b: Int): Int {
        return if (b == 0) a else gcd(b, a % b)
    }
}