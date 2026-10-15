import (
  pq "github.com/emirpasic/gods/queues/priorityqueue"
  "github.com/emirpasic/gods/utils"
)

type Twitter struct {
    count int
    // key(userid) -> slice of [2]int
    // index 0 i count, index 1 is tweetid
    tweetMap map[int][][2]int
    // key(userid) -> slice of followeeid
    followMap map[int]map[int]bool
}


func Constructor() Twitter {
    return Twitter {
        count: 0,
        tweetMap: make(map[int][][2]int),
        followMap: make(map[int]map[int]bool),
    }
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    _, e := this.followMap[userId]
    if !e {
        this.followMap[userId] = make(map[int]bool)
        this.followMap[userId][userId] = true
    }
    this.tweetMap[userId] = append(this.tweetMap[userId], [2]int{this.count, tweetId})
    this.count--
}

// TODO - minQ in Twitter struct
// minQ using container/heap for faster execution
func (this *Twitter) GetNewsFeed(userId int) []int {
    // Ordered. Most recent tweets
    var res []int
    _, e := this.followMap[userId]
    if !e {
        this.followMap[userId] = make(map[int]bool)
        this.followMap[userId][userId] = true
        return res
    }
    minQ := pq.NewWith(utils.IntComparator)
    // key - count, 
    latestTweetMap := make(map[int]int)
    for k, _ := range this.followMap[userId] {
        v, e := this.tweetMap[k]
        if !e || len(v) == 0 { continue }
        // fmt.Println(k, ". ", v)
        for i, _ := range v {
            count, tweetId := v[i][0], v[i][1]
            minQ.Enqueue(count)
            latestTweetMap[count] = tweetId
        }
    }
    if minQ.Empty() {
        return res
    }
    // fmt.Println("minQ size = ", minQ.Size(), "\n", minQ)
    tweetCount := 0
    for !minQ.Empty() {
        if tweetCount == 10 { break }
        counti, _ := minQ.Dequeue()
        count := counti.(int)
        tweetId, e := latestTweetMap[count]
        if !e { continue }
        res = append(res, tweetId)
        tweetCount++
    }
    minQ.Clear()
    return res
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if _, e := this.followMap[followerId]; !e {
        this.followMap[followerId] = make(map[int]bool)
        this.followMap[followerId][followerId] = true
    }
    this.followMap[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if _, e := this.followMap[followerId]; !e {
        return
    }
    delete(this.followMap[followerId], followeeId)
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */