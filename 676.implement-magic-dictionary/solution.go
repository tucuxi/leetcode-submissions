type MagicDictionary struct {
    words map[int][]string
}


func Constructor() MagicDictionary {
    return MagicDictionary{make(map[int][]string)}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
    words := this.words
    for _, w := range dictionary {
        words[len(w)] = append(words[len(w)], w)
    }
}


func (this *MagicDictionary) Search(searchWord string) bool {
    for _, w := range this.words[len(searchWord)] {
        mismatch := 0
        for i := range searchWord {
            if searchWord[i] != w[i] {
                mismatch++
            }
        }
        if mismatch == 1 {
            return true
        }
    }
    return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */