package _00_699




func newInteger(n int) int {
	var temp string
	for n>0{
		temp=  strconv.Itoa(n%9) + temp
		n/=9
	}
	n,_=strconv.Atoi(temp)
	return n
}
