type Trie struct {
    children map[byte]*Trie
    isTerminal bool
}

func Constructor() Trie {
    return Trie{make(map[byte]*Trie), false}    
}

func (this *Trie) Insert(word string)  {
    for i := range word {
        if _, exists := this.children[word[i]]; !exists {
            p := Constructor()
            this.children[word[i]] = &p
        }
        this = this.children[word[i]]
    }
    this.isTerminal = true
}

func (this *Trie) Search(word string) bool {
    for i := range word {
        this = this.children[word[i]]
        if this == nil {
            return false
        }
    } 
    return this.isTerminal
}

func (this *Trie) StartsWith(prefix string) bool {
    for i := range prefix {
        this = this.children[prefix[i]]
        if this == nil {
            break
        }
    } 
    return this != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */