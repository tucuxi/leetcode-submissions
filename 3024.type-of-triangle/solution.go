func triangleType(nums []int) string {
    switch a, b, c := nums[0], nums[1], nums[2]; {
    case a+b <= c || a+c <= b || b+c <= a:
        return "none"
    case a == b && b == c:
        return "equilateral"
    case a == b || a == c || b == c:
        return "isosceles"
    default:
        return "scalene"
    }
}