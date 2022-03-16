package _500_1599

import "sort"

/*
1500. 设计文件分享系统
我们需要使用一套文件分享系统来分享一个非常大的文件，该文件由 m 个从 1 到 m 编号的文件块组成。

当用户加入系统时，系统应为其注册一个独有的 ID。这个独有的 ID 应当被相应的用户使用一次，但是当用户离开系统时，其 ID 应可以被（后续新注册的用户）再次使用。

用户可以请求文件中的某个指定的文件块，系统应当返回拥有这个文件块的所有用户的 ID。如果用户收到 ID 的非空列表，就表示成功接收到请求的文件块。


实现 FileSharing 类：

FileSharing(int m) 初始化该对象，文件有 m 个文件块。
int join(int[] ownedChunks)：一个新用户加入系统，并拥有文件的一些文件块。系统应当为该用户注册一个 ID，该 ID 应是未被其他用户占用的最小正整数。返回注册的 ID。
void leave(int userID)：ID 为 userID 的用户将离开系统，你不能再从该用户提取文件块了。
int[] request(int userID, int chunkID)：ID 为 userID 的用户请求编号为 chunkID 的文件块。返回拥有这个文件块的所有用户的 ID 所构成的列表或数组，按升序排列。

*/

type FileSharing struct {
	ids []int
	maxId int
	chunk []map[int]struct{}
	// chunk  owner

}


func Constructor(m int) FileSharing {
	var f FileSharing
	for i:=0;i<=m;i++{
		f.chunk = append(f.chunk,make(map[int]struct{}))
	}
	f.maxId = 1
	return f
}


func (fs *FileSharing) Join(ownedChunks []int) int {
	var userId int
	if len(fs.ids) ==0{
		userId = fs.maxId
		fs.maxId++
	}else{
		userId = fs.ids[len(fs.ids)-1]
		fs.ids = fs.ids[:len(fs.ids)-1]
	}

	for _,chunks:=range ownedChunks{
		fs.chunk[chunks][userId] = struct{}{}
	}

	return userId
}


func (fs *FileSharing) Leave(userID int)  {
	for _,v:=range fs.chunk[1:]{
		delete(v,userID)
	}
	fs.ids = append(fs.ids,userID)
	sort.Slice(fs.ids, func(i, j int) bool {
		return fs.ids[i] >  fs.ids[j]
	})
}


func (fs *FileSharing) Request(userID int, chunkID int) []int {
	f:=fs.chunk[chunkID]
	if len(f) == 0 {
		return nil
	}
	var userIds []int
	for userId:=range f{
		userIds = append(userIds,userId)
	}

	f[userID] = struct{}{}
	sort.Ints(userIds)
	return userIds
}

