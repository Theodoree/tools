package __99

import "fmt"

// 14.最长公共前缀

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return  ""
	}

	if len(strs) == 1 {
		return  strs[0]
	}
	var minIndex int

	for i,str:=range strs{
		if len(strs[minIndex]) > len(str){
			minIndex = i
		}
	}



	var str string
	var j int
	for j < len(strs[minIndex]){

		i := 0
		for i < len(strs)-1  {
			if strs[i][j] != strs[i+1][j] {
				return str
			}
			i++
		}

		str += string(strs[minIndex][j])
		j++

	}

	return str
}

func main(){

	fmt.Println(longestCommonPrefix([]string{"a","b"}))

}