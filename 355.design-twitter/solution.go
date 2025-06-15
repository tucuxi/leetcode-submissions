type tweet struct {
    tweetId int
    userId int
}

type Twitter struct {
    tweets []*tweet
    followees map[int]map[int]bool
}

const feedLength = 10

func Constructor() Twitter {
    return Twitter{followees: map[int]map[int]bool{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.tweets = append(this.tweets, &tweet{tweetId: tweetId, userId: userId}) 
}

func (this *Twitter) GetNewsFeed(userId int) []int {
    var feed []int
    followees := this.followees[userId]
    tweets := this.tweets
    for i := len(tweets) - 1; i >= 0; i-- {
        author := tweets[i].userId
        if author == userId || followees[author] {
            feed = append(feed, tweets[i].tweetId)
            if len(feed) == feedLength {
                break
            }
        }
    }
    return feed
}

func (this *Twitter) Follow(followerId int, followeeId int)  {
    if followees := this.followees[followerId]; followees == nil {
        this.followees[followerId] = map[int]bool{followeeId: true}
    } else {
        followees[followeeId] = true
    }
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.followees[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */