func smallestBalancedIndex(nums []int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }

    ans := -1
    rightProd := 1
    rightSum := 0

    for i := len(nums) - 1; i >= 0; i-- {
        leftSum := totalSum - rightSum - nums[i]

        if leftSum == rightProd {
            ans = i 
        }

        rightSum += nums[i]

        if nums[i] == 0 {
            rightProd = 0 
        } else if rightProd > totalSum / nums[i] {
            rightProd = totalSum + 1
        } else {
            rightProd *= nums[i]
        }
    }

    return ans
}