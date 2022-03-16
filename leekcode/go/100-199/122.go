package _00_199


//122. 买卖股票的最佳时机 II

func maxProfit(prices []int) int {


	var p int
	for i:=0;i<len(prices)-1;i++{

		if prices[i] < prices[i+1]{
			p += prices[i+1] - prices[i]

		}


	}

	return  p

}