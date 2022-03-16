package _400_1499

/*
1418. 点菜展示表
给你一个数组 orders，表示客户在餐厅中完成的订单，确切地说， orders[i]=[customerNamei,tableNumberi,foodItemi] ，其中 customerNamei 是客户的姓名，tableNumberi 是客户所在餐桌的桌号，而 foodItemi 是客户点的餐品名称。
请你返回该餐厅的 点菜展示表 。在这张表中，表中第一行为标题，其第一列为餐桌桌号 “Table” ，后面每一列都是按字母顺序排列的餐品名称。接下来每一行中的项则表示每张餐桌订购的相应餐品数量，第一列应当填对应的桌号，后面依次填写下单的餐品数量。
*/
func displayTable(orders [][]string) [][]string {

	var keys []string
	var tables []int
	var m = make(map[string]map[int]int)
	m["table"] = make(map[int]int)
	for _, v := range orders {
		tableNo := atoi(v[1])
		if _, ok := m["table"][tableNo]; !ok {
			m["table"][tableNo] = 0
			tables = append(tables, tableNo)
		}

		if subm, ok := m[v[2]]; !ok {
			subm = make(map[int]int)
			subm[tableNo]++
			m[v[2]] = subm
			keys = append(keys, v[2])
		} else {
			subm[tableNo]++
		}
	}
	sort.Strings(keys)
	sort.Ints(tables)

	var result [][]string
	var lables []string
	lables = append(lables, "Table")
	lables = append(lables, keys...)

	result = append(result, lables)

	for i := 0; i < len(tables); i++ {
		var cur []string
		cur = append(cur, strconv.Itoa(tables[i]))
		for _, v := range keys {
			cur = append(cur, strconv.Itoa(m[v][tables[i]]))
		}
		result = append(result, cur)
	}

	return result

}

func atoi(str string) int {
	var cnt int
	for len(str) > 0 {
		val := int(str[0] - '0')
		str = str[1:]
		cnt = cnt*10 + val
	}
	return cnt
}
