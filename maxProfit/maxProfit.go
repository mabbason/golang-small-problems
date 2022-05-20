/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.


prices = [7,1,5,3,6,4] // 5





Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

prices = [7,6,4,3,1] // 0

setMaxProfit = 0
currProfit = 1
prices = [1,2,3,4,5,6] // buy 1  sell at 6
					B
					     S

MaxProfit = 2
currProfit = 1
prices = [2,1,2,1,0,1,2] // 2
					        B
						            S

max: 5
curr: 3
[7,1,5,3,6,4]
   B
             S

Loop until my selling day moves off the data set

 1 - generate currProfit, Val of Sell - Val of Buy
 2 - compare to currMaxProfit
		- if currProfit > maxProfit
			- overwrite maxProfit w currProfit
			- continue loop
		- else if currProfit < 0
			- move both buy day to current SellDay position
*/

package main

import "fmt"

func main() {
	fmt.Println(findMaxProfit([]int{7, 1, 5, 3, 6, 4}))    // 5
	fmt.Println(findMaxProfit([]int{1, 2, 3, 4, 5, 6, 7})) // 6
	fmt.Println(findMaxProfit([]int{2, 1, 2, 1, 0, 1, 2})) // 2
	fmt.Println(findMaxProfit([]int{7, 6, 5, 4, 3, 2}))    // 0
	fmt.Println(findMaxProfit([]int{5}))                   // 0
}

func findMaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	buy := 0
	sell := 1

	for sell < len(prices) {
		currProfit := prices[sell] - prices[buy]
		if currProfit > maxProfit {
			maxProfit = currProfit
		} else if currProfit < 0 {
			buy = sell
		}
		sell++
	}

	return maxProfit
}
