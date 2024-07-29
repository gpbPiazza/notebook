package array

// [10,7,5,8,11,2,6]

// FindMostProftiableSell receives a slice of ints representing the price of stock
// each position in the array represents one day of the year, so the index 0 of the array is day one,
// index 2 is day two and value int is the stock price in that day.
func FindMostProftiableSell(stocks []int) int {
	greatestProfit := 0
	buyPrice := stocks[0]

	for _, price := range stocks {
		potentialProfit := price - buyPrice
		if price < buyPrice {
			buyPrice = price
		} else if potentialProfit > greatestProfit {
			greatestProfit = potentialProfit
		}
	}

	return greatestProfit
}
