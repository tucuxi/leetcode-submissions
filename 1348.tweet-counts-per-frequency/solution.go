type TweetCounts struct {
    tweetTimes map[string][]int
    isSorted   map[string]bool
}


func Constructor() TweetCounts {
    return TweetCounts{
        tweetTimes: make(map[string][]int),
        isSorted:   make(map[string]bool),
    }
}


func (this *TweetCounts) RecordTweet(tweetName string, time int)  {
    this.tweetTimes[tweetName] = append(this.tweetTimes[tweetName], time)
    this.isSorted[tweetName] = false
}


func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
    if !this.isSorted[tweetName] {
        slices.Sort(this.tweetTimes[tweetName])
        this.isSorted[tweetName] = true
    }
    
    var step int
    switch freq {
    case "minute":
        step = 60
    case "hour":
        step = 3600
    case "day":
        step = 86400
    }

    numChunks := ((endTime - startTime) / step) + 1
    res := make([]int, numChunks)

    times := this.tweetTimes[tweetName]
    
    idx, _ := slices.BinarySearch(times, startTime)
    
    for i := idx; i < len(times); i++ {
        if times[i] > endTime {
            break
        }
        bucket := (times[i] - startTime) / step
        res[bucket]++
    }

    return res
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */