package _300_1399

import "sort"

/*
1348. 推文计数
请你实现一个能够支持以下两种方法的推文计数类 TweetCounts：
1. recordTweet(string tweetName, int time)
记录推文发布情况：用户 tweetName 在 time（以 秒 为单位）时刻发布了一条推文。
2. getTweetCountsPerFrequency(string freq, string tweetName, int startTime, int endTime)
返回从开始时间 startTime（以 秒 为单位）到结束时间 endTime（以 秒 为单位）内，每 分 minute，时 hour 或者 日 day （取决于 freq）内指定用户 tweetName 发布的推文总数。
freq 的值始终为 分 minute，时 hour 或者 日 day 之一，表示获取指定用户 tweetName 发布推文次数的时间间隔。
第一个时间间隔始终从 startTime 开始，因此时间间隔为 [startTime, startTime + delta*1>,  [startTime + delta*1, startTime + delta*2>, [startTime + delta*2, startTime + delta*3>, ... , [startTime + delta*i, min(startTime + delta*(i+1), endTime + 1)>，其中 i 和 delta（取决于 freq）都是非负整数。
*/

type TweetCounts struct {
	m map[string][]int
}


func Constructor() TweetCounts {
	return TweetCounts{m: make(map[string][]int)}
}


func (this *TweetCounts) RecordTweet(tweetName string, time int)  {
	this.m[tweetName] = append(this.m[tweetName],time)
}


func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	sort.Ints(this.m[tweetName])
	var unit int
	switch freq {
	case "minute":
		unit=60
	case "hour":
		unit=3600
	case "day":
		unit=3600*24
	}

	arr :=this.m[tweetName]

	var left int
	var mid int
	right:=len(arr)

	for left < right {
		mid =(left+right)>>1
		if arr[mid] < startTime {
			if left == mid {
				break
			}
			left = mid
		}else if arr [mid] > startTime {
			if right == mid {
				break
			}
			right = mid
		}else {
			break
		}
	}


	lIdx:=mid
	left = 0
	mid = 0
	right=len(arr)

	for left < right {
		mid =(left+right)>>1
		if arr[mid] < endTime {
			if left == mid {
				break
			}
			left = mid
		}else if arr [mid] > endTime {
			if right == mid {
				break
			}
			right = mid
		}else {
			break
		}
	}

	rIdx:=mid

	cnt:=(endTime-startTime)/unit+1

	var result = make([]int,cnt)
	for i:=lIdx;i<=rIdx;i++{
		if arr[i] >= startTime && arr[i] <= endTime {
			result[(arr[i]-startTime)/unit]++
		}
	}





	return result
}




