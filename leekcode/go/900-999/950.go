package _00_999

import (
	"container/list"
	"sort"
)

/*
950. 按递增顺序显示卡牌
牌组中的每张卡牌都对应有一个唯一的整数。你可以按你想要的顺序对这套卡片进行排序。
最初，这些卡牌在牌组里是正面朝下的（即，未显示状态）。
现在，重复执行以下步骤，直到显示所有卡牌为止：
从牌组顶部抽一张牌，显示它，然后将其从牌组中移出。
如果牌组中仍有牌，则将下一张处于牌组顶部的牌放在牌组的底部。
如果仍有未显示的牌，那么返回步骤 1。否则，停止行动。
返回能以递增顺序显示卡牌的牌组顺序。
答案中的第一张牌被认为处于牌堆顶部。
*/

func deckRevealedIncreasing(deck []int) []int {
	sort.Slice(deck, func(i, j int) bool {
		return deck[i] < deck[j]
	})

	l := list.New()
	for i := 0; i < len(deck); i++ {
		l.PushBack(i)
	}

	a := make([]int, 0, len(deck))
	for l.Len() > 0 {
		a = append(a, l.Front().Value.(int))
		l.Remove(l.Front())

		if l.Len() > 0 {
			v := l.Front().Value.(int)
			l.Remove(l.Front())
			l.PushBack(v)
		}
	}

	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[a[i]] = deck[i]
	}

	return b
}
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	var arr []int
	for i:=len(deck)-1;i>=0;i--{
		arr = append([]int{deck[i]},arr...)

		if len(arr) >=2 {
			n := arr[len(arr)-1]
			copy(arr[2:],arr[1:])
			arr[1] = n
		}


	}

	return arr
}

