object Solution {
    def removeElement(nums: Array[Int], `val`: Int): Int = {
        var k = 0
        for (num <- nums) {
            if (num != `val`) {
                nums(k) = num
                k += 1
            }
        }
        k
    }
}