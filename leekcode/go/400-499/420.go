package _00_499



/*
420. 强密码检验器
如果一个密码满足下述所有条件，则认为这个密码是强密码：
由至少 6 个，至多 20 个字符组成。
至少包含 一个小写 字母，一个大写 字母，和 一个数字 。
同一字符 不能 连续出现三次 (比如 "...aaa..." 是不允许的, 但是 "...aa...a..." 如果满足其他条件也可以算是强密码)。
给你一个字符串 password ，返回 将 password 修改到满足强密码条件需要的最少修改步数。如果 password 已经是强密码，则返回 0 。
在一步修改操作中，你可以：
插入一个字符到 password ，
从 password 中删除一个字符，或
用另一个字符来替换 password 中的某个字符。
*/
func strongPasswordChecker(password string) int {
	line:=password
	hasTypes:=0
	hasLower,hasUpper,hasNum:=false,false,false
	needDeleteNum, needReplaceNum :=0,0
	n:=len(line)
	counts:=make([]int,0)
	i:=0
	for i<n{
		count:=0
		temp:=i
		for temp<n&&line[temp]==line[i]{
			temp++
			count++
		}
		counts=append(counts,count)
		if !hasNum&&line[i]>='0'&&line[i]<='9'{
			hasNum=true
			hasTypes++
		}
		if !hasLower&&line[i]>='a'&&line[i]<='z'{
			hasLower=true
			hasTypes++
		}
		if !hasUpper&&line[i]>='A'&&line[i]<='Z'{
			hasUpper=true
			hasTypes++
		}
		i=temp
	}

	if n<=3{
		return 6-n
	}else if n<6{
		return max(3-hasTypes,6-n)
	}else if n>20{
		needDeleteNum=n-20
	}
	needDele2:=needDeleteNum
	for i := 0; i < 2; i++ {
		for j, num := range counts {
			if needDeleteNum>=i+1{
				if num > 2 && num%3 == i {
					counts[j] = num - i-1
					needDeleteNum=needDeleteNum-i-1
				}
			}else{
				break
			}
		}
	}
	if needDeleteNum>0{
		for i,num:=range counts{
			if needDeleteNum>0{
				if num>2{
					if num-2>=needDeleteNum{
						counts[i]=counts[i]-needDeleteNum
						needDeleteNum=0
					}else{
						needDeleteNum=needDeleteNum-(num-2)
						counts[i]=2
					}
				}
			}else{
				break
			}
		}
	}
	if needDeleteNum>0{
		return 3-hasTypes+needDele2
	}
	for _,num:=range counts{
		needReplaceNum+=num/3
	}
	return max(3-hasTypes,needReplaceNum)+needDele2
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}