package _go

/*
面试题 16.02. 单词频率
设计一个方法，找出任意指定单词在一本书中的出现频率。
你的实现应该支持如下操作：
WordsFrequency(book)构造函数，参数为字符串数组构成的一本书
get(word)查询指定单词在书中出现的频率
示例：
WordsFrequency wordsFrequency = new WordsFrequency({"i", "have", "an", "apple", "he", "have", "a", "pen"});
wordsFrequency.get("you"); //返回0，"you"没有出现过
wordsFrequency.get("have"); //返回2，"have"出现2次
wordsFrequency.get("an"); //返回1
wordsFrequency.get("apple"); //返回1
wordsFrequency.get("pen"); //返回1
*/

type WordsFrequency struct { // 用个锤子前缀树,大力出奇迹,最大长度为10其实应该用前缀树,但是hash表不超时,用个锤子trie
	m map[string]int
}

func Constructor(book []string) WordsFrequency {
	var w WordsFrequency
	w.m = make(map[string]int, len(book)/2)
	for _, v := range book {
		w.m[v]++
	}
	return w
}

func (this *WordsFrequency) Get(word string) int {
	return this.m[word]
}
