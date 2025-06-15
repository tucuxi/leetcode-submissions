func maskPII(s string) string {
    var b strings.Builder
    
    if a := strings.IndexByte(s, '@'); a >= 0 {
        email := strings.ToLower(s)
        b.WriteByte(email[0])
        b.WriteString("*****")
        b.WriteString(email[a-1:])
    } else {
        phone := make([]byte, 0, 13)       
        for i := range s {
            if s[i] >= '0' && s[i] <= '9' {
                phone = append(phone, s[i])
            }
        }
        if len(phone) > 10 {
            b.WriteByte('+')
            b.WriteString(strings.Repeat("*", len(phone)-10))
            b.WriteByte('-')
        }
        b.WriteString("***-***-")
        b.WriteString(string(phone[len(phone)-4:]))
    }
    return b.String()
}