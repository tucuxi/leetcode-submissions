func simplifyPath(path string) string {
    components := []string{}
    for _, s := range strings.Split(path, "/") {
        switch s {
            case "", ".":
                break
            case "..":
                if len(components) > 0 {
                    components = components[:len(components) - 1]
                }
            default:
                components = append(components, s)
        }
    }
    return "/" + strings.Join(components, "/")
}