type AuthenticationManager struct {
    ttl int
    tokens map[string]int
}


func Constructor(timeToLive int) AuthenticationManager {
    return AuthenticationManager{
        ttl: timeToLive,
        tokens: make(map[string]int),
    }
}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
    this.tokens[tokenId] = currentTime + this.ttl
}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
    if this.tokens[tokenId] > currentTime {
        this.tokens[tokenId] = currentTime + this.ttl
    }
}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
    res := 0
    for tokenId, expiryTime := range this.tokens {
        if expiryTime <= currentTime {
            delete(this.tokens, tokenId)
        } else {
            res++
        }
    }
    return res
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */