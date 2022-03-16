package _500_1599

/*
1507. 转变日期格式
给你一个字符串 date ，它的格式为 Day Month Year ，其中：

Day 是集合 {"1st", "2nd", "3rd", "4th", ..., "30th", "31st"} 中的一个元素。
Month 是集合 {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"} 中的一个元素。
Year 的范围在  [1900, 2100] 之间。
请你将字符串转变为 YYYY-MM-DD 的格式，其中：

YYYY 表示 4 位的年份。
MM 表示 2 位的月份。
DD 表示 2 位的天数。
*/

func reformatDate(date string) string {
	var (
		idx    int
		result [10]byte
	)

	result[4] = '-'
	result[7] = '-'

	if date[1] >= '0' && date[1] <= '9' {
		result[8] = date[0]
		result[9] = date[1]
	} else {
		result[8] = '0'
		result[9] = date[0]
	}

	for date[idx] != ' ' {
		idx++
	}
	idx++

	result[5] = '0'
	switch date[idx : idx+3] {
	case "Jan":
		result[6] = '1'
	case "Feb":
		result[6] = '2'
	case "Mar":
		result[6] = '3'
	case "Apr":
		result[6] = '4'
	case "May":
		result[6] = '5'
	case "Jun":
		result[6] = '6'
	case "Jul":
		result[6] = '7'
	case "Aug":
		result[6] = '8'
	case "Sep":
		result[6] = '9'
	case "Oct":
		result[5] = '1'
		result[6] = '0'
	case "Nov":
		result[5] = '1'
		result[6] = '1'
	case "Dec":
		result[5] = '1'
		result[6] = '2'
	}

	idx += 3 + 1

	var resultIdx = 0
	for idx < len(date) {
		result[resultIdx] = date[idx]
		resultIdx++
		idx++
	}

	return string(result[:])
}
