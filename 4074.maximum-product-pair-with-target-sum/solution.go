func maxProductPair(nums []int, target int) []int {
    h := make(map[int]int)
    bestI, bestJ := -1, -1
    product := math.MinInt

    for i := range nums {
        wanted := target - nums[i]
        if wanted != nums[i] && nums[i] * wanted > product {
            if index, exists := h[wanted]; exists {
                bestI = i
                bestJ = index
                product = nums[i] * wanted
            }
        }
        h[nums[i]] = i
    }

    if bestI == -1 {
        return []int{-1, -1}
    }
    if nums[bestI] < nums[bestJ] {
        return []int{bestJ, bestI}
    }
    return []int{bestI, bestJ}
}