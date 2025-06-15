object Solution {
    def twoSum(nums: Array[Int], target: Int): Array[Int] = {
        val withIndex = nums.zipWithIndex
        val map = withIndex.toMap
        withIndex.collectFirst {
            case (value, index) if map.get(target - value).exists(_ != index) =>
                Array(index, map(target - value))
        }.getOrElse(Array(-1, -1))
    }            
}