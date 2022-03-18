package _00_699


/*
609. 在系统中查找重复文件
给你一个目录信息列表 paths ，包括目录路径，以及该目录中的所有文件及其内容，请你按路径返回文件系统中的所有重复文件。答案可按 任意顺序 返回。
*/
func findDuplicate(paths []string) [][]string {
	var m = make(map[string][]string)
	for _, v := range paths {
		arr := strings.Split(v, " ")
		for _, file := range arr[1:] {
			idx := strings.Index(file, "(")
			absFileName := arr[0] + "/" + file[:idx]
			context := file[idx+1 : len(file)-1]
			m[context] = append(m[context], absFileName)
		}
	}
	var arr [][]string
	for _, v := range m {
		if len(v) > 1 {
			arr = append(arr, v)
		}
	}
	return arr
}
