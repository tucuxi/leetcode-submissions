const maxNum = 50

func largestInteger(nums []int, k int) int {
    var h [maxNum + 1]int

    for i := k; i <= len(nums); i++ {
        s := make(map[int]bool)
        for j := i - k; j < i; j++ {
            s[nums[j]] = true
        }
        for num := range s {
            h[num]++
        }
    }

    for num := maxNum; num >= 0; num-- {
        if h[num] == 1 {
            return num
        }
    }

    return -1
}