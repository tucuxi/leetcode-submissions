func decodeMessage(key string, message string) string {
    k := [26]byte{}
    j := byte('a')
    for i := range key {
        if c := key[i]; c != ' ' {
            if k[c - 'a'] == 0 {
                k[c - 'a'] = j
                j++
                if j > 'z' {
                    break
                }
            }
        } 
    }
    for j <= 'z' {
        k[j - 'a'] = j
        j++
    }
    var sb strings.Builder
    for i := range message {
        if message[i] == ' ' {
            sb.WriteByte(' ')
        } else {
            sb.WriteByte(k[message[i] - 'a'])
        }
    }
    return sb.String()
}