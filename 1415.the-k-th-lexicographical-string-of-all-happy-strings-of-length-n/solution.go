func getHappyString(n int, k int) string {
    var res string 
    c := 0
    
    var rec func([]byte)
    rec = func(pre []byte) {
        if c == k {
            return
        }
        if len(pre) == n {
            c++
            if c == k {
                res = string(pre)
            }
            return
        }
        for ch := byte('a'); ch <= byte('c'); ch++ {
            if ch != pre[len(pre) - 1] {
                rec(append(pre, ch))
            }
        }
        
    }
    
    for ch := byte('a'); ch <= byte('c'); ch++ {
        rec([]byte{ch})
    }
    
    return res
}