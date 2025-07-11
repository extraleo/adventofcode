package hot100

func maxProfit(prices []int) int {
    ans, minPrice := -1,prices[0]
		for _, p := range(prices){
			ans = max(ans, p-minPrice)
			minPrice = min(p, minPrice)
		}
		return ans
}