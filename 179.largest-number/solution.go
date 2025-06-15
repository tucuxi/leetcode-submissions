func largestNumber(nums []int) string {
    s := make([]string, len(nums))
    for i := range nums {
        s[i] = strconv.Itoa(nums[i])
    }
    sort.Slice(s, func(i, j int) bool {
        return s[j] + s[i] < s[i] + s[j]
    })
    if s[0] == "0" {
        return "0"
    }
    return strings.Join(s, "")
}