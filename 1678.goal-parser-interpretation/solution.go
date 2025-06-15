func interpret(command string) string {
    var sb strings.Builder
    for i := 0; i < len(command); {
        if command[i] == 'G' {
            sb.WriteByte('G')
            i++
        } else if command[i+1] == ')' {
            sb.WriteByte('o')
            i += 2
        } else {
            sb.WriteString("al")
            i += 4
        }
    }
    return sb.String()
}