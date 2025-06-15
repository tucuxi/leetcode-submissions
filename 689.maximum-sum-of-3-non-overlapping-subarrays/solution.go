func maxSumOfThreeSubarrays(nums []int, k int) []int {
    prefixSum := func() []int {
        res := make([]int, len(nums) + 1)
        for i, num := range nums {
            res[i + 1] = res[i] + num
        }
        return res
    }()

    leftMaxIndex := func() []int {
        res := make([]int, len(nums))
        currentMaxSum := prefixSum[k] - prefixSum[0]
        for i := k; i < len(nums); i++ {
            if d := prefixSum[i + 1] - prefixSum[i + 1 - k]; d > currentMaxSum {
                currentMaxSum = d
                res[i] = i + 1 - k
            } else {
                res[i] = res[i - 1]
            }
        }
        return res
    }()

    rightMaxIndex := func() []int {
        n := len(nums)
        res := make([]int, n)
        res[n - k] = n - k
        currentMaxSum := prefixSum[n] - prefixSum[n - k]
        for i := n - k - 1; i >= 0; i-- {
            if d := prefixSum[i + k] - prefixSum[i]; d >= currentMaxSum {
                currentMaxSum = d
                res[i] = i
            } else {
                res[i] = res[i + 1]
            }
        }
        return res
    }()

    return func() []int {
        var res []int
        maxSum := 0
        for i := k; i <= len(nums) - 2 * k; i++ {
            leftIndex := leftMaxIndex[i - 1]
            rightIndex := rightMaxIndex[i + k]
            totalSum := prefixSum[i + k] - prefixSum[i] +
                prefixSum[leftIndex + k] - prefixSum[leftIndex] +
                prefixSum[rightIndex + k] - prefixSum[rightIndex]
            if totalSum > maxSum {
                maxSum = totalSum
                res = []int{leftIndex, i, rightIndex}
            }
        }
        return res
    }()
}