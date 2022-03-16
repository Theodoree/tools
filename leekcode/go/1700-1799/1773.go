package _700_1799

/*
1773. 统计匹配检索规则的物品数量
给你一个数组 items ，其中 items[i] = [typei, colori, namei] ，描述第 i 件物品的类型、颜色以及名称。
另给你一条由两个字符串 ruleKey 和 ruleValue 表示的检索规则。
如果第 i 件物品能满足下述条件之一，则认为该物品与给定的检索规则 匹配 ：
ruleKey == "type" 且 ruleValue == typei 。
ruleKey == "color" 且 ruleValue == colori 。
ruleKey == "name" 且 ruleValue == namei 。
统计并返回 匹配检索规则的物品数量 。
*/
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var count int
	for _, v := range items {
		switch ruleKey {
		case "type":
			if v[0] == ruleValue {
				count++
			}
		case "color":
			if v[1] == ruleValue {
				count++
			}
		case "name":
			if v[2] == ruleValue {
				count++
			}
		}
	}
	return count
}
