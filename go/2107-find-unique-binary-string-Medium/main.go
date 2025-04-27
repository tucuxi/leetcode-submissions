func findDifferentBinaryString(nums []string) string {
    var sb strings.Builder
    for i := range nums {
        sb.WriteByte(nums[i][i] ^ 1)
    }
    return sb.String()
}