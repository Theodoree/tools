package _600_1699

/*
1600. 王位继承顺序
一个王国里住着国王、他的孩子们、他的孙子们等等。每一个时间点，这个家庭里有人出生也有人死亡。
这个王国有一个明确规定的王位继承顺序，第一继承人总是国王自己。我们定义递归函数 Successor(x, curOrder) ，给定一个人 x 和当前的继承顺序，该函数返回 x 的下一继承人。
*/

type ThroneInheritance struct {
	king  string
	m     map[string][]string
	depth map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	var t = ThroneInheritance{
		king:  kingName,
		m:     make(map[string][]string),
		depth: make(map[string]bool),
	}

	return t
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.m[parentName] = append(t.m[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.depth[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	var arr []string
	dfs(t.m, t.depth, t.king, &arr)
	return arr
}

func dfs(m map[string][]string, depth map[string]bool, cur string, arr *[]string) {
	if !depth[cur] {
		*arr = append(*arr, cur)
	}
	for _, v := range m[cur] {
		dfs(m, depth, v, arr)
	}
}

