type LUPrefix struct {
    prefix int
    uploads []bool
}


func Constructor(n int) LUPrefix {
    return LUPrefix{0, make([]bool, n + 1)}
}


func (this *LUPrefix) Upload(video int)  {
    if !this.uploads[video] {
        this.uploads[video] = true
        if video == this.prefix + 1 {
            for video < len(this.uploads) && this.uploads[video] {
                this.prefix++
                video++
            }
        }
    }
}


func (this *LUPrefix) Longest() int {
    return this.prefix
}


/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */