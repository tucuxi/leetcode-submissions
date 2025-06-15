func isNumber(s string) bool {
    matched, _ := regexp.MatchString(`^(\+|-)?(\d+(\.\d*)?|\.\d+)([eE](\+|-)?\d+)?$`, s)
    return matched 
}