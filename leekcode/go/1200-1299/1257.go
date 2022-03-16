package _200_1299


/*
1257. 最小公共区域
给你一些区域列表 regions ，每个列表的第一个区域都包含这个列表内所有其他区域。
很自然地，如果区域 X 包含区域 Y ，那么区域 X  比区域 Y 大。
给定两个区域 region1 和 region2 ，找到同时包含这两个区域的 最小 区域。
如果区域列表中 r1 包含 r2 和 r3 ，那么数据保证 r2 不会包含 r3 。
数据同样保证最小公共区域一定存在。
*/
func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	var m = map[string]string{}
	for idx:=range regions{
		for _,v:=range regions[idx][1:]{
			m[v] = regions[idx][0]
		}
	}

	var hs =map[string]struct{}{}

	for region1 != ""{
		hs[region1] = struct{}{}
		region1 = m[region1]
	}
	for {
		if _,ok:=hs[region2];ok{
			break
		}
		region2 = m[region2]
	}


	return region2
}




