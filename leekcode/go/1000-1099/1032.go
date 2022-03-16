package _000_1099

/*
1032. 字符流
按下述要求实现 StreamChecker 类：

StreamChecker(words)：构造函数，用给定的字词初始化数据结构。
query(letter)：如果存在某些 k >= 1，可以用查询的最后 k个字符（按从旧到新顺序，包括刚刚查询的字母）拼写出给定字词表中的某一字词时，返回 true。否则，返回 false。


示例：

StreamChecker streamChecker = new StreamChecker(["cd","f","kl"]); // 初始化字典
streamChecker.query('a');          // 返回 false
streamChecker.query('b');          // 返回 false
streamChecker.query('c');          // 返回 false
streamChecker.query('d');          // 返回 true，因为 'cd' 在字词表中
streamChecker.query('e');          // 返回 false
streamChecker.query('f');          // 返回 true，因为 'f' 在字词表中
streamChecker.query('g');          // 返回 false
streamChecker.query('h');          // 返回 false
streamChecker.query('i');          // 返回 false
streamChecker.query('j');          // 返回 false
streamChecker.query('k');          // 返回 false
streamChecker.query('l');          // 返回 true，因为 'kl' 在字词表中。
*/
type StreamChecker struct {
	root  tree
	idx   int
	str   []byte
	count int
}

type tree struct {
	treeArr [26]*tree
	match   bool
}

func (t *tree) Insert(str string) {
	cur := t
	for idx := len(str) - 1; idx >= 0; idx-- {
		v := str[idx]
		if cur.treeArr[v-'a'] == nil {
			cur.treeArr[v-'a'] = &tree{}
		}
		cur = cur.treeArr[v-'a']
	}
	cur.match = true
}

func Constructor(words []string) StreamChecker {
	var stream StreamChecker
	var max int
	for _, v := range words {
		stream.root.Insert(v)
		if len(v) > max {
			max = len(v)
		}
	}
	stream.str = make([]byte, 0, max)
	stream.count = max
	return stream
}

func (s *StreamChecker) Query(letter byte) bool {
	letter -= 'a'
	if len(s.str) < s.count {
		s.str = append(s.str, letter)
	} else {
		s.str[s.idx] = letter
	}
	var ok bool
	cur := &s.root
	idx := s.idx

	s.idx++
	if s.idx >= s.count {
		s.idx = 0
	}

	for i := 0; i < len(s.str); i++ {
		if cur == nil {
			break
		}

		bytes := s.str[idx]
		idx--
		if idx < 0 {
			idx = len(s.str) - 1
		}

		cur = cur.treeArr[bytes]
		if cur != nil && cur.match {
			ok = true
			break
		}

	}

	return ok
}
