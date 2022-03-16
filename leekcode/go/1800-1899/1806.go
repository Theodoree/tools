package _800_1899




func reinitializePermutation(n int) int {
	/*
	   因为每个数字都有固定的移动轨迹，所以当它回到原来位置时，所有元素都回到了原来的位置上
	   例如n=6，操作历史
	   [0, 1, 2, 3, 4, 5]
	   [0, 3, 1, 4, 2, 5]
	   [0, 4, 3, 2, 1, 5]
	   [0, 2, 4, 1, 3, 5]
	   [0, 1, 2, 3, 4, 5]
	   可以看到：
	   1. 开头和结尾元素位置肯定不变；
	   2.其他位置都是从这个2->3->5->4->2->3->5->4路径循环
	       下标0: 1->1->1->1->1，都在位置1
	       下标1: 2->3->5->4->2，一开始在2，移动到位置3，5，4，然后又回到2
	       下标2: 3->5->4->2->3，一开始在3，移动到位置5，4，2，然后又回到3
	       下标3: 4->2->3->5->4，一开始在4，移动到位置2，3，5，然后又回到4
	       下标4: 5->4->2->3->5，一开始在5，移动到位置4，2，3，然后又回到5
	       下标5: 6->6->6->6->6，都在位置6
	   3.一个数字如位置2的，从2位置开始， 它会跑到3，之后它就会按位置3开始的元素相同的运动轨迹开始跑。因为在同一位置处到达下一位置的方法是唯一的，也就是说，除了开头的数字和结尾的数字以外，所有元素的路径都是重叠的。故所以判断下标1是否回到了自己的位置。故下标1: 2->3->5->4->2这个需要经过5步
	*/
	var per=make([]int,n)
	var next=make([]int,n)
	for i:=0;i<n;i++{
		per[i]=i
	}
	for i:=0;i<n;i++{
		if i%2==0{
			next[i]=per[i/2]
		}else{
			next[i]=per[n/2+(i-1)/2]
		}
	}
	var res int
	var idx=1
	for {
		res++
		idx=next[idx]
		if idx==1{
			break
		}
	}
	return res
}
