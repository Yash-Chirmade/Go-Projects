package solution

//121. Best Time to Buy and Sell Stock 
//{7,1,5,3,6,4}
func MaxProfit(prices []int) int {
	minPrice := prices[0]
	max := 0
	for i := 0; i <= len(prices)-1; i++ {

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if max < prices[i]-minPrice {
			max = prices[i] - minPrice
		}

	}
	return max
}
