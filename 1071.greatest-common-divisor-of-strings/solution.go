func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }
    gcd := euclid(len(str1), len(str2))
    return str1[:gcd]
}

func euclid(a, b int) int {
    for a % b > 0 {
        a, b = b, a % b
    }
    return b
}