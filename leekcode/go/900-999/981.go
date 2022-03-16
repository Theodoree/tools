package _00_999

/*
981. 基于时间的键值存储
设计一个基于时间的键值数据结构，该结构可以在不同时间戳存储对应同一个键的多个值，并针对特定时间戳检索键对应的值。
实现 TimeMap 类：
TimeMap() 初始化数据结构对象
void set(String key, String value, int timestamp) 存储键 key、值 value，以及给定的时间戳 timestamp。
String get(String key, int timestamp)
返回先前调用 set(key, value, timestamp_prev) 所存储的值，其中 timestamp_prev <= timestamp 。
如果有多个这样的值，则返回对应最大的  timestamp_prev 的那个值。
如果没有值，则返回空字符串（""）。
*/
type val struct {
	value string
	ts    int
}

type TimeMap struct {
	m map[string]*[]val
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string]*[]val)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	arr, ok := t.m[key]
	if !ok {
		arr = &[]val{}
		t.m[key] = arr
	}
	*arr = append(*arr, val{
		value: value,
		ts:    timestamp,
	})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	_arr, ok := t.m[key]
	if !ok || (*_arr)[0].ts > timestamp {
		return ""
	}

	arr := *_arr
	if len(arr) == 0 {
		return ""
	}
	left := 0
	right := len(arr)
	mid := 0
	for left < right {
		mid = (left + right) >> 1
		if arr[mid].ts > timestamp {
			if right == mid {
				break
			}
			right = mid
		} else if arr[mid].ts < timestamp {
			if left == mid {
				break
			}
			left = mid
		} else {
			break
		}
	}

	return arr[mid].value

}
