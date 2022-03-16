package _100_1199



/*
1166. 设计文件系统
你需要设计一个文件系统，你可以创建新的路径并将它们与不同的值关联。
路径的格式是一个或多个连接在一起的字符串，形式为： / ，后面跟着一个或多个小写英文字母。例如， " /leetcode" 和 "/leetcode/problems" 是有效路径，而空字符串 "" 和 "/" 不是。
实现 FileSystem 类:
bool createPath(string path, int value) 创建一个新的 path ，并在可能的情况下关联一个 value ，然后返回 true 。如果路径已经存在或其父路径不存在，则返回 false 。
 int get(string path) 返回与 path 关联的值，如果路径不存在则返回 -1 。
*/

type FileSystem struct {
	system map[string]*FileSystem
	value  int
	exits  bool
}


func Constructor() FileSystem {
	return FileSystem{system: make(map[string]*FileSystem)}
}


func (f *FileSystem) CreatePath(path string, value int) bool {
	subPath:=strings.Split(path,"/")
	subF:=f
	for idx,v:=range subPath{
		if v == ""{
			continue
		}
		if _,ok:=subF.system[v];!ok{
			if idx != len(subPath)-1 { // 父路径不存在
				return false
			}
			subF.system[v] = &FileSystem{system: make(map[string]*FileSystem)}
		}
		subF = subF.system[v]
	}

	if subF.exits { // 该路径已存在
		return false
	}

	subF.exits = true
	subF.value = value
	return true
}


func (f *FileSystem) Get(path string) int {
	subPath:=strings.Split(path,"/")
	subF:=f

	for _,v:=range subPath{
		if v == ""{
			continue
		}
		if _,ok:=subF.system[v];!ok{
			return -1
		}
		subF = subF.system[v]
	}

	return subF.value
}

