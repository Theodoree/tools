package __99

import "fmt"

//22.括号生成 回溯大法好
/*
var result []string

func generateParenthesis(n int) []string {

	var array []string
	parentheses("", n, n)
	array = result
	result = []string{}
	return array
}

func parentheses(sublist string, leftNum int, rightNum int) {
	if leftNum == 0 && rightNum == 0 {
		result = append(result, sublist)
	}

	if rightNum > leftNum {
		parentheses(sublist+")", leftNum, rightNum-1)
	}
	if leftNum > 0 {
		parentheses(sublist+"(", leftNum-1, rightNum)
	}

}


*/
/*
func generateParenthesis(n int) []string {
	var ret []string

	backtrack(&ret, "", n, n)
	return ret

}

func backtrack(ret *[]string, ans string, left, right int) {
	if left == 0 && right == 0 {
		*ret = append(*ret, ans)
	}

	if left > 0 {
		backtrack(ret, ans+"(", left-1, right)
	}
	if right > left {
		backtrack(ret, ans+")", left, right-1)
	}
}
*/

func generateParenthesis(n int) []string {
	ret := []string{}

	backtrack(&ret, "", 0, 0, n)
	return ret

}

func backtrack(ret *[]string, ans string, left, right , n int){
	if len(ans) == 2*n {
		*ret = append(*ret, ans)
		return
	}
	if left < n {
		backtrack(ret, ans+"(", left+1, right, n)
	}
	if right < left {
		backtrack(ret, ans+")", left, right+1, n)
	}
}
func main() {

	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))

}
