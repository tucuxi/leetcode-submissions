func compareVersion(version1 string, version2 string) int {
    v1 := strings.Split(version1, ".")
    v2 := strings.Split(version2, ".")
    for i := len(v2) - len(v1); i > 0; i-- {
        v1 = append(v1, "0")
    } 
    for i := len(v1) - len(v2); i > 0; i-- {
        v2 = append(v2, "0")
    }
    for i := range v1 {
        n1, _ := strconv.Atoi(v1[i])
        n2, _ := strconv.Atoi(v2[i])
        if n1 < n2 {
            return -1
        }
        if n1 > n2 {
            return 1
        }
    }
    return 0
}