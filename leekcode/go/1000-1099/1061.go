package _000_1099

/*
1061. 按字典序排列最小的等效字符串
给出长度相同的两个字符串s1 和 s2 ，还有一个字符串 baseStr 。
其中  s1[i] 和 s2[i]  是一组等价字符。
举个例子，如果 s1 = "abc" 且 s2 = "cde"，那么就有 'a' == 'c', 'b' == 'd', 'c' == 'e'。
等价字符遵循任何等价关系的一般规则：
 自反性 ：'a' == 'a'
 对称性 ：'a' == 'b' 则必定有 'b' == 'a'
 传递性 ：'a' == 'b' 且 'b' == 'c' 就表明 'a' == 'c'
例如， s1 = "abc" 和 s2 = "cde" 的等价信息和之前的例子一样，那么 baseStr = "eed" , "acd" 或 "aab"，这三个字符串都是等价的，而 "aab" 是 baseStr 的按字典序最小的等价字符串
利用 s1 和 s2 的等价信息，找出并返回 baseStr 的按字典序排列最小的等价字符串。
*/

type ufs struct {
	roots []int
	cnt   int
}

func newUfs(n int) *ufs {
	roots := make([]int, n)
	for i := range roots {
		roots[i] = i
	}
	return &ufs{roots, n}
}
func (v *ufs) Find(x int) int {
	root := x
	for root != v.roots[root] {
		root = v.roots[root]
	}
	for root != v.roots[x] {
		cur := v.roots[x]
		v.roots[x] = root
		x = cur
	}
	return root
}

func (v *ufs) Union(x, y int) {
	xRoot, yRoot := v.Find(x), v.Find(y)
	if xRoot == yRoot {
		return
	}
	if yRoot < xRoot {
		v.cnt--
		v.roots[xRoot] = yRoot
	}
}

func (v *ufs) Count() int {
	return v.cnt
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {

	ufs := newUfs(26)
	for i := 0; i < len(s1); i++ {
		ufs.Union(int(s1[i]-'a'), int(s2[i]-'a'))
		ufs.Union(int(s2[i]-'a'), int(s1[i]-'a'))
	}

	for i := 0; i < 26; i++ {
		ufs.Find(i)
	}

	var result []byte
	for i := 0; i < len(baseStr); i++ {
		n := byte(ufs.roots[baseStr[i]-'a'] + 'a')
		result = append(result, n)
	}

	return string(result)
}
