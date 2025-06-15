func validIPAddress(queryIP string) string {
    if validIPv4(queryIP) {
        return "IPv4"
    }
    if validIPv6(queryIP) {
        return "IPv6"
    }
    return "Neither"
}

func validIPv4(queryIP string) bool {
    re := regexp.MustCompile(`^(\d{1,3})\.(\d{1,3}).(\d{1,3}).(\d{1,3})$`)
    s := re.FindStringSubmatch(queryIP)
    if len(s) != 5 {
        return false
    }
    for i := 1; i < len(s); i++ {
        if s[i][0] == '0' && len(s[i]) > 1 {
            return false
        }
        b, _ := strconv.Atoi(s[i])
        if b > 255 {
            return false
        }
    }
    return true
}

func validIPv6(queryIP string) bool {
    matched, _ := regexp.MatchString(`^[a-fA-F0-9]{1,4}(:[a-fA-F0-9]{1,4}){7}$`, queryIP)
    return matched
}